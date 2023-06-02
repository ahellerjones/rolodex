<script>
  import Auth from "./lib/Auth.svelte";
  import Contacts from "./lib/Contacts.svelte";

  let loggedIn = false;
</script>

{#if loggedIn}
  <Contacts />
{:else}
  <Auth />
{/if}
