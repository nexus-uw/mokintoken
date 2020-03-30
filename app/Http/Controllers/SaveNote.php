<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Note;
use DateTime;


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
        $expiresViews = $request->input('expiresViews');

        $expires = new \DateTime();
        $expires->add(new \DateInterval(\sprintf('PT%dH', $request->input('expiresHours'))));

        $note = new Note;
        $note->encryptedText = $request->input('encryptedText');
        $note->expiry = $expires;
        // todo add max view count and/or expiry datetime
        $note->id = uniqid();
        $note->viewCount = 0;
        $note->save();

        return response()->json([
            'id'=>$note->id
        ]);
    }
}
