<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Note;


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
        $note = Note::where('id', $id)->first();
        // todo check for expiry datetime or view count
        // if expired + view too much. delete + return info (either not found or deleted)
        // else inc view count by 1
        return view('decrypt',['encryptedText' => $note->encryptedText]);
    }
}
