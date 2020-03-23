<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
    <head>

        <title>Laravel</title>
        @include('style')

        <style>
            #share{
                display: none;
            }
        </style>
    </head>
    <body class="code ">
    <div class="center measure-wide">
    <h1 class="f-subheadline-ns f1  lh-solid mt0 mb3">MokinToken</h1>
    <h2 class=" mt0 fw3">secure note sharing for the 2077 normalization</h2>
    <div>
    <a>home</a>
    <a>about</a>
    </div>
    <div id="newNote">

        <form  id="newNoteForm" class="pb7">
          <div class="pa3 bg-white black br2">
            <textarea id="text" class="w-100" autocapitalize="none" autocomplete="off" autofocus maxlength="50000" rows=
            "25"></textarea>
          </div>
          <div>
            <button id="submit" class="grow w-100 ba bw2 ">encrypt and submit</button>
          </div>
        </form>
    </div>
    <div id="share">
        <h1>note successfully encrypted and ready to share</h1>
        <div>either share this link <a id="shareLink" target="_blank">holla</a>or this QR code</div>
        <div> <canvas id="shareQR"></canvas></div>
        <button> generate another note? </button>
    </div>

        <script type="module" src="/index.js" ></script>
    </body>
</html>
