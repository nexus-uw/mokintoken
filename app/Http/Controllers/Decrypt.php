<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Note;
use DateTime;

class Decrypt extends Controller
{
    /**
     * Handle the incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function __invoke($id)
    {
        $note = Note::where(
          'id', $id
          )
          ->first();

        // if expired + view too much. delete + return info (either not found or deleted)
        // else inc view count by 1
        if( (!is_null($note)) && ($note->viewCount >= $note->expiresViews || new \DateTime($note->expiry) < new \DateTime("now") ) ){
          Note::where('id', $id)->delete();
          $note = NULL;
        }


        if (is_null($note)){
          return view('noteDoesNotExist');
        } else {
          Note::where('id', $id)
            ->increment('viewCount');
          return view('decrypt',['encryptedText' => $note->encryptedText]);
        }

    }
}
