<script>
  import { LoadConfigs } from '../wailsjs/go/main/App.js'
  import { onMount } from 'svelte'
  import SaveConfigs from "./SaveConfigs.svelte";

  let configLoaded = false;
  let email, token, baseURL

  function saveConfigs() {
    SaveConfigs(email, token, baseURL)
  }

  onMount(async () => {
    LoadConfigs().then(result => {
      if (result) {
        configLoaded = true
        email = result.Email
        token = result.Token
        baseURL = result.BaseURL
      }
    })
  })
</script>

<main>
    {#if !configLoaded}
        <SaveConfigs/>
    {:else}
        <div>Your email is {email}</div>
    {/if}
</main>