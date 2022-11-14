<script>
  import {SaveConfigs} from '../wailsjs/go/main/App.js'

  let email, token, baseURL

  function saveConfigs() {
    SaveConfigs(email, token, baseURL)
  }
</script>

<main>
  <div>{email}</div>
  <form on:submit|preventDefault={saveConfigs}>
    <label>
      BaseURL: <input type="text" bind:value={baseURL} >
    </label>
    <label>
      Email: <input type="text" bind:value={email}>
    </label>
    <label>
      Token: <input type="password" bind:value={token}>
    </label>
    <button type="submit">Submit</button>
  </form>
</main>