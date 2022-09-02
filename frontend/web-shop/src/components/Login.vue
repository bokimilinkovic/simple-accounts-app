<template>
    <form @submit.prevent="login" class="size">
    <!-- Email input -->
        <div class="form-group row">
        <label for="staticEmail" class="col-sm-2 col-form-label">Email</label>
        <div class="col-sm-10">
        <input type="text" class="form-control-plaintext" id="staticEmail" v-model="email">
        </div>
    </div>
    <div class="form-group row">
        <label for="inputPassword" class="col-sm-2 col-form-label">Password</label>
        <div class="col-sm-10">
        <input type="password" class="form-control" id="inputPassword" placeholder="Password" v-model="password">
        </div>
    </div>

    <!-- Submit button -->
    <button type="submit" class="btn btn-primary btn-block mb-4 ">Sign in</button>

    <!-- Register buttons -->
    <div class="text-center">
        <p>Not a member? <a href="#!">Register</a></p>
        <p>or sign up with:</p>
        </div>
    </form>
</template>

<script>
import { ref } from "vue"
import { useRouter } from "vue-router";
import {useStore} from "vuex";

import axiosInstance from "../axios/index"

export default {
    name: "Login",
    setup(){
        const email = ref("")
        const password = ref("")
        const token = ref("")
        const store = useStore()
        const router = useRouter()


        async function login() {
            console.log(email.value)
            console.log(password.value)
            if(email.value == '' || password.value=='') {
                alert("Email or pass can't be empty")
                return 
            }
            try{
                const resp = await axiosInstance.post("/signin",
                {
                    email: email.value,
                    password: password.value
                })
                if(resp.data.token){
                    store.commit('addApiToken',`Bearer ${resp.data.token}`)
                    localStorage.setItem('api-token', resp.data.token)
                }
                router.push('/movies')
            }catch(err) {
                console.log(err)
            }
        }

        return{
            email,
            password,
            login,
            token
        }
    }
}
</script>

<style>

</style>