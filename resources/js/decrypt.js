import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'

export class Decrypt {
  loadNonceAndKey() {
    let nonce, key

    if (!window.location.hash || window.location.hash === '') {
      console.error('decryption key not present in hash...')
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

  constructor() {
    this.text = document.getElementById('text')

    this.encryptedtext = document.getElementById('encryptedtext')
    this.decryptedtext = document.getElementById('decryptedtext')

    const { nonce, key } = this.loadNonceAndKey()

    // debugging only, remove once ready
    this.decryptedtext.value = naclUutil.encodeUTF8(
      nacl.secretbox.open(
        naclUutil.decodeBase64(this.encryptedtext.value),
        nonce,
        key
      )
    )
  }
}
