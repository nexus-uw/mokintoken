import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'

console.log("ass " + nacl + naclUutil + Base64)
// https://stackoverflow.com/questions/51562781/service-worker-fetch-event-is-not-firing
self.addEventListener('activate', function (event) {
  console.log('Claiming control')
  return self.clients.claim()
})
self.addEventListener("fetch", (event) => {

  // Regular requests not related to Web Share Target.
  if (event.request.method !== "POST" || new URL(event.request.url).pathname != "/share-target/") {
    event.respondWith(fetch(event.request))
    return
  }

  // Requests related to Web Share Target.
  event.respondWith(
    (async () => {

      const formData = await event.request.formData()
      const nonce = nacl.randomBytes(nacl.secretbox.nonceLength)
      const key = nacl.randomBytes(nacl.secretbox.keyLength)
      // encrypt note or image
      // send to fetch api
      const imgBase64 = null // todo - base64
      const f = await fetch('/api/save-note', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          encryptedText: naclUutil.encodeBase64(
            nacl.secretbox(
              naclUutil.decodeUTF8(formData.get('text')),
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
          // code is hard
          expiresHours: "4",
          expiresViews: "1"
        })
      })
      const { id } = await f.json()

      const hash = Base64.encode(
        JSON.stringify({
          nonce: naclUutil.encodeBase64(nonce),
          key: naclUutil.encodeBase64(key)
        })
      )
      // respond with saved note page
      return Response.redirect(`/noteSaved#id=${id}&hash=${hash}`, 303)
    })(),
  )
})
