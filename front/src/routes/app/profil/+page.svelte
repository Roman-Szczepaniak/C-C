<script>
  let user = {
    name: 'John Doe',
    email: 'john.doe@example.com'
  };
</script>

<main class="container mx-auto p-4">
  <h1 class="text-3xl font-bold mb-4">Profil Utilisateur</h1>
  <div class="p-4 border rounded-md">
    <p><strong>Nom:</strong> {user.name}</p>
    <p><strong>Email:</strong> {user.email}</p>
  </div>
</main>
