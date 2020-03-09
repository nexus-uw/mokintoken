const { nacl } = window; // should already be attached to the window

class Main {
  setNonceAndKey(nonce, key) {
    // set in url hash
  }

  loadNonceAndKey() {
    let nonce, key;

    if (!window.location.hash || window.location.hash === '') {
      key = nacl.randomBytes(nacl.secretbox.keyLength);
      nonce = nacl.randomBytes(nacl.secretbox.nonceLength);
    } else {
      // read from url hash
    }

    return {
      nonce,
      key
    };
  }

  constructor() {
    this.text = document.getElementById('text');

    this.encryptedtext = document.getElementById('encryptedtext');
    this.decryptedtext = document.getElementById('decryptedtext');

    const { nonce, key } = this.loadNonceAndKey();

    // for debugging purposes only, todo: remove once ready
    this.key = document.getElementById('key');
    this.nonce = document.getElementById('nonce');
    this.nonce.value = nacl.util.encodeBase64(nonce);
    this.key.value = nacl.util.encodeBase64(key);

    this.text.addEventListener('change', e => {
      this.setNonceAndKey(nonce, key);

      const newValue = e.target.value;

      this.encryptedtext.value = nacl.util.encodeBase64(
        window.nacl.secretbox(nacl.util.decodeUTF8(newValue), nonce, key)
      );

      // debugging only, remove once ready
      this.decryptedtext.value = nacl.util.encodeUTF8(
        window.nacl.secretbox.open(
          nacl.util.decodeBase64(this.encryptedtext.value),
          nonce,
          key
        )
      );
    });
  }
}

new Main();
