<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
  <head>
    <title>MokinToken: secure note sharing for the 2077 normalization</title>
    <meta name="description" content="selfhosted e2e encrypted note sharing webapp"/>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="stylesheet" href="/tachyons.min.css" />
    <meta charset="utf-8" />

    <style>
      #share{
        display: none;
      }
    </style>
  </head>
  <body class="code ">
    <noscript>JS is required to do anything with this site...</noscript>
    <div class="center measure-wide">
      <h1 class="f-subheadline-ns f1  lh-solid mt0 mb3">MokinToken</h1>
      <h2 class=" mt0 fw3">secure note sharing for the 2077 normalization</h2>
      <a href="/">home</a>
      <a href="/about">about</a>

      <div class="container">
        @yield('content')
      </div>
    </div>
    <footer class="pv4 mv4 bt center tc">
        <a href="http://mokinan4qvxi4ragyzgkewrmnnqslkcdglk6v5zruknwnnuvv2lu5uad.onion" > TOR </a> | 2020 - CURRENT YEAR <a href="https://unlicense.org/">Unlicensed</a> | Another project by <a href="https://ramsay.xyz/?ref=mokintoken">Simon Ramsay</a> | <a href="https://github.com/nexus-uw/mokintoken">CODE</a>

    </footer>
    <script type="module" src="/index.js" ></script>
 </body>
</html>
