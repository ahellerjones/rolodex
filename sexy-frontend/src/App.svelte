<script>
  import Auth from "./components/Auth.svelte";
  import Contacts from "./components/Contacts.svelte";
  import { auth } from "./lib/stores";

  let loggedIn;
  auth.subscribe((val) => {
    loggedIn = val;
  });
</script>

{#if loggedIn}
  <Contacts />
{:else}
  <Auth />
{/if}
