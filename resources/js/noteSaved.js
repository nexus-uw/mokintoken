import QRCode from 'qrcode'


export class NoteSaved {

  constructor() {
    const hash = window.location.hash.slice(1)
    const id = new URLSearchParams(window.location.search).get('id')

    const clearnetShareLink = `${document.querySelector('meta[name="clearnet"]').getAttribute('content')}/decrypt/${id}#${hash}`
    document.getElementById('clearnetShareLink').href = clearnetShareLink

    const dakrnetShareLink = `${document.querySelector('meta[name="darknet"]').getAttribute('content')}/decrypt/${id}#${hash}`
    document.getElementById('dakrnetShareLink').href = dakrnetShareLink

    Promise.all([
      QRCode.toCanvas(document.getElementById('clearnetShareQR'), clearnetShareLink),
      QRCode.toCanvas(document.getElementById('darknetShareQR'), dakrnetShareLink)
    ])
  }
}