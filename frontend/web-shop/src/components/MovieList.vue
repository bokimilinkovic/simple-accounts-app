<template>
  <div class="album py-5 bg-light">
        <div class="container">
          <div class="row">
            <div class="col-md-4" v-for="movie in movies" :key="movie.ID">
              <div class="card mb-4 shadow-sm">
                <img class="card-img-top" :src="movie.CoverURL" data-holder-rendered="true" style="height: 225px; width: 100%; display: block;">
                <div class="card-body">
                  <p class="card-text">{{movie.Description}}</p>
                  <div class="d-flex justify-content-between align-items-center">
                    <div class="btn-group">
                      <button type="button" class="btn btn-sm btn-outline-secondary">View</button>
                      <button type="button" class="btn btn-sm btn-outline-secondary">Edit</button>
                    </div>
                    <large class="text-muted">{{movie.Price}} e</large>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
</template>

<script>
import {  ref } from 'vue'
import axiosInstance from "../axios/index"


export default {
    name: "MovieList",
    setup(){
        
        // let movies = computed(function(){
        //   return store.state.movies
        // })

        let movies = ref({})

        async function fetchMovies(){
            try {
                const resp = await axiosInstance.get('/movies')
                if(resp.status == 200) {
                  movies.value = resp.data
                }

                console.log(movies.value)
            } catch (error) {
                console.log(error)
            }
        }
        fetchMovies()

        return{
            movies,
            // fetchMovies
        }
    },
}
</script>

<style>

</style>