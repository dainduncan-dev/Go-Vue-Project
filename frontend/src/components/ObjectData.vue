<template>
  <div class="object-container">
    <div class="object-data">
      <h1>Name: {{ name }}</h1>
    </div>
    <div class="object-data">
      <h1>Description: {{ description }}</h1>
    </div>
    <div class="object-data">
      <h1>Attributes: {{ attributes }}</h1>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      name: null,
      description: null,
      attributes: null
    }
  },

  // Fetches posts when the component is created.
  created() {
    axios.get(`http://localhost:1323/get`)
    .then(response => {
      // JSON responses are automatically parsed.
      this.name = response.data.name;
      this.description = response.data.description;
      this.attributes = response.data.attributes;
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.object-container {
  width: 100%;
  height: 100vh fit-content;
  background-color: #171718;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: left;
  border-radius: 5px;
  
}

.object-data {
  padding: 60px;
  width: 100%;
  color: white;
  margin-left: 60px;
}

h1 {
  margin: 40px 0 0;
  text-align: left;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
