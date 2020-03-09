<?hh

require __DIR__."/vendor/hh_autoload.php";

<<__EntryPoint>>
function main(): noreturn {
echo (

      <head>
        <title>mokintoken project</title>
        <meta charset="utf-8"/>
      </head>
);

$request_uri = explode('?', $_SERVER['REQUEST_URI'], 2);

// todo: use proper router...
switch ($request_uri[0]) {
    // Home page
    case '/':
        require __DIR__.'/src/home.hack';
        break;
    case '/about':
        require __DIR__.'/src/about.hack';
        break;
    case (preg_match('/\/note\/*/',$request_uri[0])?true:false):
      require __DIR__.'/src/decrypt.hack';
      break;
    default:
        header('HTTP/1.0 404 Not Found');
        require __DIR__.'/src/404.hack';
        break;
}

echo (
  <div>
    <footer>
      foot stuff?
    </footer>
    <script src="https://unpkg.com/tweetnacl"/>
    <script src="https://unpkg.com/tweetnacl-util"/>
    <script src="https://unpkg.com/js-base64"/>
    <script src="/assets/index.js" />
  </div>
);
  exit(0);
}
