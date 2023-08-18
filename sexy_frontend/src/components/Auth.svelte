<script>
  import { POST_LOGIN } from "../lib/constants";
  import { auth } from "../lib/stores";

  let loading = false;
  let username;
  let password;

  const onClick = (loginOrSetup) => {
    if (username && password) {
      loading = true;

      const body = {
        username,
        password,
        loginOrSetup,
      };

      fetch(POST_LOGIN, {
        method: "POST",
        body: JSON.stringify(body),
      })
        .then((res) => {
          console.log(res);
          loading = false;
        })
        .catch((err) => {
          console.log(err);
          loading = false;
          auth.set(true);
        });
    }
  };
</script>

<main class=" h-screen w-screen flex justify-center items-center">
  {#if loading}
    <div>loading</div>
  {:else}
    <div class=" w-96 xl:w-1/3 p-4 rounded-md border-2 space-y-4 flex flex-col">
      <h1 class=" font-bold text-xl">Authenticate</h1>
      <input
        type="text"
        bind:value={username}
        placeholder="username"
        class=" p-2 border rounded"
      />
      <input
        type="text"
        bind:value={password}
        placeholder="password"
        class=" p-2 border rounded"
      />
      <div class=" flex space-x-4">
        <button
          class=" bg-slate-200 p-3 rounded-sm flex-1 hover:bg-slate-300"
          on:click={() => onClick(true)}>Log In</button
        >
        <button
          class=" bg-slate-200 p-3 rounded-sm flex-1 hover:bg-slate-300"
          on:click={() => onClick(false)}>Sign Up</button
        >
      </div>
    </div>
  {/if}
</main>
