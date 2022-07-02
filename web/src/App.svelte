<script>
  import { Router, Route } from "svelte-routing";
  import LoginPage from "./lib/LoginPage.svelte";

  let url = "";
</script>

<main>
  <Router {url}>
    <Route path="/"><LoginPage /></Route>
  </Router>
</main>

<style>
</style>
