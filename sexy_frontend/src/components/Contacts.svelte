<script>
  import { onMount } from "svelte";
  import { GET_CONTACTS } from "../lib/constants";
  import Header from "./Header.svelte";
  import PlusIcon from "../assets/PlusIcon.svelte";
  import Contact from "./Contact.svelte";

  let contacts;
  let loading;
  let editing = false;

  // fetch contacts list
  onMount(() => {
    loading = true;

    fetch(GET_CONTACTS)
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        loading = false;
      });
  });
</script>

<main>
  {#if loading}
    <div>loading</div>
  {:else}
    <Header />
    <div
      class=" flex flex-col justify-center items-center p-10 mt-24 space-y-6"
    >
      {#each [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1] as yo}
        <Contact />
      {/each}
    </div>
    <div class=" fixed right-10 bottom-10">
      <button class=" p-4 bg-green-200 hover:bg-green-300 rounded-full"
        ><PlusIcon /></button
      >
    </div>
  {/if}
</main>
