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

        if ($expiresViews < 1){
          return response()->json([
            'message' => 'invalid expiresViews'], 404);
        }

        $expires = new \DateTime();
        $expires->add(new \DateInterval(\sprintf('PT%dH', $request->input('expiresHours'))));

        $note = new Note;
        $note->encryptedText = $request->input('encryptedText');
        $note->expiry = $expires;
        $note->id = uniqid();
        $note->viewCount = 0;
        $note->expiresViews = $expiresViews;
        $note->save();

        return response()->json([
            'id'=>$note->id
        ]);
    }
}
