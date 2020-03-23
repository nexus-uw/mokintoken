@extends('layouts.app')

@section('content')

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
  <div class="db"> <canvas id="shareQR" style="display:block; margin: auto;"></canvas></div>
</div>

@endsection
