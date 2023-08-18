<script>
  import { auth, user } from "../lib/stores";

  let username;
  user.subscribe((val) => {
    username = val.username;
  });

  const logout = () => {
    auth.set(false);
  };

  const deleteAccount = () => {
    console.log("delete");
  };
</script>

<header
  class=" border-b-2 p-6 fixed top-0 w-full bg-white flex justify-between items-center"
>
  <h1>{username}</h1>
  <div class=" space-x-4">
    <button
      on:click={() => logout()}
      class=" p-4 rounded bg-slate-200 hover:bg-slate-300">Log Out</button
    >
    <button
      on:click={() => deleteAccount()}
      class=" p-4 rounded text-red-900 bg-red-200 hover:bg-red-300"
      >Delete Account</button
    >
  </div>
</header>
