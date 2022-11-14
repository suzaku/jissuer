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

  import Dropzone from "svelte-file-dropzone";

  let files = {
      accepted: [],
      rejected: []
  };

  function handleFilesSelect(e) {
      const { acceptedFiles, fileRejections } = e.detail;
      files.accepted = [...files.accepted, ...acceptedFiles];
      files.rejected = [...files.rejected, ...fileRejections];
  }
</script>

<main>
    {#if !configLoaded}
        <h1>Hello</h1>
        <SaveConfigs/>
    {:else}
        <div>Hello {email}</div>
        <Dropzone on:drop={handleFilesSelect} />
        <ol>
            {#each files.accepted as item}
                <li>{item.name}</li>
            {/each}
        </ol>
    {/if}
</main>