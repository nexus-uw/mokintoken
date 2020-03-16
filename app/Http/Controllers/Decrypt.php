<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

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
        return view('decrypt',['encryptedText' => 'aIyYebVLEOhA37SIO6vTclISsaa4S9jibyZMAHkQQD/w2g4tLrdldg8=']);
    }
}
