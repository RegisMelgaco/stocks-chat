<svelte:head>
    <script src="https://accounts.google.com/gsi/client" async defer></script>
    <script>
        function onLogin(resp) {
            document.cookie = "credential="+resp.credential;
        }
    </script>
</svelte:head>

<div
    id="g_id_onload"
    data-client_id="698154937532-gj4c6ofq2t0q03ffv7vbimuf0cfrc1g8.apps.googleusercontent.com"
    data-login_uri="http://localhost:3000"
    data-auto_prompt="false"
    data-callback="onLogin"
/>
<div
    class="g_id_signin"
    data-type="standard"
    data-theme="outline"
    data-text="sign_in_with"
    data-logo_alignment="left"
/>
