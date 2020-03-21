<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <title>Laravel</title>

        <!-- Fonts -->
        <link href="https://fonts.googleapis.com/css?family=Nunito:200,600" rel="stylesheet">

        <!-- Styles -->
        <style>
            html, body {
                background-color: #fff;
                color: #636b6f;
                font-family: 'Nunito', sans-serif;
                font-weight: 200;
                height: 100vh;
                margin: 0;
            }

            .full-height {
                height: 100vh;
            }

            .flex-center {
                align-items: center;
                display: flex;
                justify-content: center;
            }

            .position-ref {
                position: relative;
            }

            .top-right {
                position: absolute;
                right: 10px;
                top: 18px;
            }

            .content {
                text-align: center;
            }

            .title {
                font-size: 84px;
            }

            .links > a {
                color: #636b6f;
                padding: 0 25px;
                font-size: 13px;
                font-weight: 600;
                letter-spacing: .1rem;
                text-decoration: none;
                text-transform: uppercase;
            }

            .m-b-md {
                margin-bottom: 30px;
            }

            .share{
                display: 'none';
            }
        </style>
    </head>
    <body>
    <div>
    <div id="newNote">
        <form method="POST" id="newNoteForm">
          <label>key</label><input id="key" readonly="true" />
          <label>nonce</label><input id="nonce" readonly="true" />
          <label>text</label><textarea id="text" ></textarea>
          <label>encrypted text</label><textarea id="encryptedtext" name="encryptedtext" ></textarea>
          <label>decryptedtext text</label><textarea id="decryptedtext" disabled="true" ></textarea>
          <button type="submit" >submit</button>
        </form>
    </div>
    <div id="share">
        <h1>note successfully encrypted and ready to share</h1>
        <div>either share this link <a id="shareLink">holla</a></div>
        <div>or this QR code <img id="shareQR"></img></div>
    </div>

        <script src="/js/app.js" ></script>
    </body>
</html>
