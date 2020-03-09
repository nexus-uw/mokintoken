require __DIR__."/../vendor/hh_autoload.php";

<<__EntryPoint>>
function intro_examples_avoid_xss_main(): noreturn {
  echo (
    <html>
      <head>
        <title>mokintoken project</title>
        <meta charset="utf-8"/>
      </head>
      <body>
        <div>
          <label>key</label><input id="key" readonly="true" />
          <label>nonce</label><input id="nonce" readonly="true" />
          <label>text</label><textarea id="text" />
          <label>encrypted text</label><textarea id="encryptedtext" disabled="true" />
          <label>decryptedtext text</label><textarea id="decryptedtext" disabled="true" />

        </div>
        <script src="https://unpkg.com/tweetnacl"/>
        <script src="https://unpkg.com/tweetnacl-util"/>

        <script src="index.js" />
      </body>
    </html>
  );
  exit(0);
}