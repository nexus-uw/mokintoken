import nacl from 'tweetnacl'
import naclUutil from 'tweetnacl-util'
import { Base64 } from 'js-base64'

console.log("ass" + nacl + naclUutil + Base64)


self.addEventListener("fetch", (event) => {
  // Regular requests not related to Web Share Target.
  if (event.request.method !== "POST" || new URL(event.request.url).pathname != "/assets/share-target") {
    event.respondWith(fetch(event.request))
    return
  }

  // Requests related to Web Share Target.
  event.respondWith(
    (async () => {
      const formData = await event.request.formData()
      // encrypt note or image
      // send to fetch api
      // respond with saved note page
      return Response.redirect('/noteSaved#todoshitinthehash', 303)
    })(),
  )
})
