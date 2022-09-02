<template>
  <header>
        <nav class="navbar navbar-expand-lg navbar-light bg-white">
    <div class="container-fluid" style="
      background-image: url('https://mdbcdn.b-cdn.net/img/new/slides/041.webp');
      height: 200px;
    ">
      <button
        class="navbar-toggler"
        type="button"
        data-mdb-toggle="collapse"
        data-mdb-target="#navbarExample01"
        aria-controls="navbarExample01"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <i class="fas fa-bars"></i>
      </button>
      <div class="collapse navbar-collapse" id="navbarExample01">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item active">
            <a class="nav-link text-danger" aria-current="page" href="#">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link  text-danger" href="/login">Login</a>
          </li>
          <li class="nav-item">
            <a class="nav-link  text-danger" href="/counter">Counter</a>
          </li>
          <li class="nav-item">
            <a class="nav-link  text-danger" href="/movies">Movies</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">About</a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
    </header>
</template>

<script>
export default {
    name: "Header"
}
</script>

<style>

</style>