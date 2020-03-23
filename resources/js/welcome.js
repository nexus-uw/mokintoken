import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'
import QRCode from 'qrcode'

export class Welcome {
  setNonceAndKey(nonce, key) {
    // set in url hash
    window.location.hash = Base64.encode(
      JSON.stringify({
        nonce: naclUutil.encodeBase64(nonce),
        key: naclUutil.encodeBase64(key)
      })
    )
  }

  loadNonceAndKey() {
    let nonce, key

    if (!window.location.hash || window.location.hash === '') {
      key = nacl.randomBytes(nacl.secretbox.keyLength)
      nonce = nacl.randomBytes(nacl.secretbox.nonceLength)
    } else {
      // read from url hash
      const json = JSON.parse(Base64.decode(window.location.hash))
      nonce = naclUutil.decodeBase64(json.nonce)
      key = naclUutil.decodeBase64(json.key)
    }

    return {
      nonce,
      key
    }
  }

  saveNote() {}

  constructor() {
    this.text = document.getElementById('text')

    const { nonce, key } = this.loadNonceAndKey()

    // for debugging purposes only, todo: remove once ready

    this.text.addEventListener('change', e => {
      this.setNonceAndKey(nonce, key)

      const newValue = e.target.value

      this.encryptedtext = naclUutil.encodeBase64(
        nacl.secretbox(naclUutil.decodeUTF8(newValue), nonce, key)
      )
    })

    const newNoteForm = document.getElementById('newNoteForm')
    newNoteForm.addEventListener('submit', async e => {
      e.preventDefault()

      const f = await fetch('/api/save-note', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          encryptedText: this.encryptedtext
        })
      })

      const { id } = await f.json()

      document.getElementById('newNote').style.display = 'none'
      document.getElementById('share').style.display = 'initial'
      const shareLink = `${location.protocol}//${location.host}/decrypt/${id}${location.hash}`
      document.getElementById('shareLink').href = shareLink

      await QRCode.toCanvas(document.getElementById('shareQR'), shareLink)
      return false
    })
  }
}
