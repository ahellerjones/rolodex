<script>
  import CalendarIcon from "../assets/CalendarIcon.svelte";
  import HomeIcon from "../assets/HomeIcon.svelte";
  import MailIcon from "../assets/MailIcon.svelte";
  import PhoneIcon from "../assets/PhoneIcon.svelte";
  import TrashIcon from "../assets/TrashIcon.svelte";

  let confirmDelete = false;
  let editing = false;

  let contact = {
    name: "jim bob",
    address: "1433 Lombard Street",
    phoneNumber: "(862) 621-8674",
    email: "jim.bob@gmail.com",
    birthday: "09-20-2000",
  };
</script>

<div
  class=" w-96 xl:w-1/3 p-4 flex flex-col border-2 rounded-md space-y-4 text-lg"
>
  <div class=" flex justify-between space-x-4">
    {#if confirmDelete}
      <button
        on:click={() => {
          confirmDelete = false;
        }}
        class=" p-4 flex-1 rounded bg-slate-200 hover:bg-slate-300"
        >cancel</button
      >
      <button
        class=" p-4 flex-1 rounded text-red-900 bg-red-200 hover:bg-red-300"
        >delete</button
      >
    {:else}
      <h1 class=" font-bold text-2xl">{contact.name}</h1>
      <button
        on:click={() => {
          confirmDelete = true;
        }}><TrashIcon /></button
      >
    {/if}
  </div>
  <div class=" flex space-x-10 hover:bg-slate-100 p-2 rounded cursor-text">
    <HomeIcon />
    <span>
      {contact.address}
    </span>
  </div>
  <div class=" flex space-x-10 hover:bg-slate-100 p-2 rounded">
    <PhoneIcon />
    <span>
      {contact.phoneNumber}
    </span>
  </div>
  <div class=" flex space-x-10 hover:bg-slate-100 p-2 rounded">
    <MailIcon />
    <span>
      {contact.email}
    </span>
  </div>
  <div class=" flex space-x-10 hover:bg-slate-100 p-2 rounded">
    <CalendarIcon />
    <span>
      {contact.birthday}
    </span>
  </div>
</div>
