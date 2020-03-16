<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Note;

class SaveNote extends Controller
{
    /**
     * Handle the incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function __invoke(Request $request)
    {

        $note = new Note;
        $note->encryptedText = $request->input('encryptedText');
        // todo add max view count and/or expiry datetime
        $note->id = uniqid();
        $note->save();

        return response()->json([
            'id'=>$note->id
        ]);
    }
}
