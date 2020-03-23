import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'
import QRCode from 'qrcode'

export class Welcome {
  loadNonceAndKey() {
    // always generate new nonce
    const nonce = nacl.randomBytes(nacl.secretbox.nonceLength)

    let key
    if (!window.location.hash || window.location.hash === '') {
      key = nacl.randomBytes(nacl.secretbox.keyLength)
    } else {
      // read from url hash
      const json = JSON.parse(Base64.decode(window.location.hash))
      key = naclUutil.decodeBase64(json.key)
    }

    // set in url hash
    window.location.hash = Base64.encode(
      JSON.stringify({
        nonce: naclUutil.encodeBase64(nonce),
        key: naclUutil.encodeBase64(key)
      })
    )

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
            )
          })
        })

        const { id } = await f.json()

        document.getElementById('newNote').style.display = 'none'
        document.getElementById('share').style.display = 'initial'
        const shareLink = `${location.protocol}//${location.host}/decrypt/${id}${location.hash}`
        document.getElementById('shareLink').href = shareLink

        await QRCode.toCanvas(document.getElementById('shareQR'), shareLink)
      } catch (e) {
        console.error('failed to save note', e)
      }
      return false
    })
  }
}
