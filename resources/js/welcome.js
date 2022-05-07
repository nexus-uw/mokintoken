import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'
import QRCode from 'qrcode'

export class Welcome {
  loadNonceAndKey() {
    window.location.hash = '' // clear out old key
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
