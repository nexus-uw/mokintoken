import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'
import QRCode from 'qrcode'

async function getImgBase64(ele) {
  if (ele.files.length == 0) {
    return null //img is optional
  }
  const reader = new FileReader()
  reader.readAsDataURL(ele.files[0])
  return new Promise((resolve, reject) => {
    reader.onloadend = () => resolve(reader.result)
    reader.onerror = (e) => reject(e)
  })
}

export class Welcome {
  loadNonceAndKey() {
    //window.location.hash = '' // clear out old key
    // always generate new nonce
    const nonce = nacl.randomBytes(nacl.secretbox.nonceLength)
    const key = nacl.randomBytes(nacl.secretbox.keyLength)

    return {
      nonce,
      key
    }
  }

  constructor() {
    const { nonce, key } = this.loadNonceAndKey()

    const newNoteForm = document.getElementById('newNoteForm')
    newNoteForm.addEventListener('submit', async e => {
      e.preventDefault()
      document.getElementById('submit').disabled = true
      try {

        const imgBase64 = await getImgBase64(document.getElementById('img'))

        const f = await fetch('/api/save-note', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            encryptedText: naclUutil.encodeBase64(
              nacl.secretbox(
                naclUutil.decodeUTF8(document.getElementById('text').value),
                nonce,
                key
              )
            ),
            encryptedImg: imgBase64 != null ? naclUutil.encodeBase64(
              nacl.secretbox(
                naclUutil.decodeUTF8(imgBase64),
                nonce,
                key
              )
            ) : null,
            expiresHours: document.getElementById('expiresHours').value,
            expiresViews: document.getElementById('expiresViews').value
          })
        })

        const { id } = await f.json()
        document.getElementById('newNote').style.display = 'none'
        document.getElementById('share').style.display = 'initial'
        const hash = Base64.encode(
          JSON.stringify({
            nonce: naclUutil.encodeBase64(nonce),
            key: naclUutil.encodeBase64(key)
          })
        )

        const clearnetShareLink = `${document.querySelector('meta[name="clearnet"]').getAttribute('content')}/decrypt/${id}#${hash}`
        document.getElementById('clearnetShareLink').href = clearnetShareLink

        const dakrnetShareLink = `${document.querySelector('meta[name="darknet"]').getAttribute('content')}/decrypt/${id}#${hash}`
        document.getElementById('dakrnetShareLink').href = dakrnetShareLink

        await Promise.all([
          QRCode.toCanvas(document.getElementById('clearnetShareQR'), clearnetShareLink),
          QRCode.toCanvas(document.getElementById('darknetShareQR'), dakrnetShareLink)
        ])
      } catch (e) {
        alert('failed to save note: ' + e)
      }
      document.getElementById('submit').disabled = false
      return false
    })
  }
}
