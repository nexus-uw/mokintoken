require __DIR__."/../vendor/hh_autoload.php";


  switch ( $_SERVER['REQUEST_METHOD']){
    case 'GET':
echo (

        <div>
        <form method="POST">
          <label>key</label><input id="key" readonly="true" />
          <label>nonce</label><input id="nonce" readonly="true" />
          <label>text</label><textarea id="text" />
          <label>encrypted text</label><textarea id="encryptedtext" name="encryptedtext"  />
          <label>decryptedtext text</label><textarea id="decryptedtext" disabled="true" />
          <button type="submit">submit</button>
        </form>
        </div>

  );
    break;

    case 'POST':
    // this should only be ajax endpoint.
echo (

        <div>
        POSTED {$_POST['encryptedtext']}
        </div>


  );
    break;
    default:
      http_response_code(415);
      break;
  }

