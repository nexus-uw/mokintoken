@extends('layouts.app')


@section('content')
<main>


<h1>about</h1>
<h2>the name MokinToken</h2>
<p>2 things that were on my desk when I came up with the repo name
a mokin usb c hub and a security token</p>


<h2>purpose</h2>
<p>I wanted a selfhosted secure note sharing site for myself to send text between my computers.
I have found this to a recurring issue when I have had to login a site on a someone elses computer but my password is stored in a password manager on my phone and is extremely long and complex </p>
<p>Existing sites either wanted too much infrastructure or did not provide a docker container. Some additional people created docker containers for the projects, but they were not being updated with the latest application code.</p>

<p>also i wanted something to do that was quick and easy</p>


<h2>security</h2>
<div>
<ul>
<li>
 client side encrypted: <a href="https://tweetnacl.js.org/#/">tweetnacl</a>
</li>
<li>
 secret is stored in the <a href="https://en.wikipedia.org/wiki/Fragment_identifier">URL hash</a> (never sent to the server)
</li>
<li>
 source: <a href="https://github.com/nexus-uw/mokintoken">nexus-uw/mokintoken</a>
</li>
<li>
easliy <a href="https://hub.docker.com/r/nexusuw/mokintoken">self hosted</a>
</li>
</ul>
</div>
<h2>issues:</h2>
<div>
<ul>
<li>
3rd party audit -> NOPE
</li>
<li>
trustworthy -> NOPE
</li>
<li>
proffessional -> NOPE
</li>
<li>
<a href="https://security.stackexchange.com/questions/128581/hosting-company-advised-us-to-avoid-php-for-security-reasons-are-they-right/128587">PHP...</a>
</li>
</ul>
</div>

<h2>alternative</h2>
<p>
<a href="https://github.com/awesome-selfhosted/awesome-selfhosted#custom-communication-systems">self host your own e2e encrypted chat application for just yourself and your devices</a>
</p>
<p>
<a href="https://github.com/awesome-selfhosted/awesome-selfhosted#note-taking-and-editors">other self hosted note apps</a>
</p>
</main>

@endsection
