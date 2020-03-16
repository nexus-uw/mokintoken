const nacl = require('tweetnacl')
nacl.util = require('tweetnacl-util')
import { Base64 } from 'js-base64'


class Main {
    setNonceAndKey(nonce, key) {
        // set in url hash
        window.location.hash = Base64.encode(
            JSON.stringify({
                nonce: nacl.util.encodeBase64(nonce),
                key: nacl.util.encodeBase64(key)
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
            nonce = nacl.util.decodeBase64(json.nonce)
            key = nacl.util.decodeBase64(json.key)
        }

        return {
            nonce,
            key
        }
    }

    saveNote() {

    }

    constructor() {
        this.text = document.getElementById('text')

        this.encryptedtext = document.getElementById('encryptedtext')
        this.decryptedtext = document.getElementById('decryptedtext')

        const { nonce, key } = this.loadNonceAndKey()

        // for debugging purposes only, todo: remove once ready
        this.key = document.getElementById('key')
        this.nonce = document.getElementById('nonce')
        this.nonce.value = nacl.util.encodeBase64(nonce)
        this.key.value = nacl.util.encodeBase64(key)

        this.text.addEventListener('change', e => {
            this.setNonceAndKey(nonce, key)

            const newValue = e.target.value

            this.encryptedtext.value = nacl.util.encodeBase64(
                nacl.secretbox(nacl.util.decodeUTF8(newValue), nonce, key)
            )

            // debugging only, remove once ready
            this.decryptedtext.value = nacl.util.encodeUTF8(
                nacl.secretbox.open(
                    nacl.util.decodeBase64(this.encryptedtext.value),
                    nonce,
                    key
                )
            )
        })

        const newNoteForm = document.getElementById('newNoteForm')
        newNoteForm.addEventListener('submit', async e => {

            e.preventDefault()

            await fetch('/api/save-note', {

                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    encryptedText: this.encryptedtext.value
                })
            })

            // todo: show share info with the ID generated from the api


            return false

        })
    }
}


class Decrypt {
    loadNonceAndKey() {
        let nonce, key

        if (!window.location.hash || window.location.hash === '') {
            console.error('decryption key not present in hash...')
        } else {
            // read from url hash
            const json = JSON.parse(Base64.decode(window.location.hash))
            nonce = nacl.util.decodeBase64(json.nonce)
            key = nacl.util.decodeBase64(json.key)
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

        // for debugging purposes only, todo: remove once ready
        this.key = document.getElementById('key')
        this.nonce = document.getElementById('nonce')
        this.nonce.value = nacl.util.encodeBase64(nonce)
        this.key.value = nacl.util.encodeBase64(key)

        // debugging only, remove once ready
        this.decryptedtext.value = nacl.util.encodeUTF8(
            nacl.secretbox.open(
                nacl.util.decodeBase64(this.encryptedtext.value),
                nonce,
                key
            )
        )


    }
}

if (location.pathname === '/') {
    new Main()
}

if (location.pathname.startsWith('/decrypt')) {
    new Decrypt()
}
