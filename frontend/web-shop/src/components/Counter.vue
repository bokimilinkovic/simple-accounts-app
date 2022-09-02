<template>
    <button @click="increment">
    Count is: {{ state.count }}, double is: {{ state.double }}
  </button>
</template>

<script>
import {reactive, computed} from "vue"

export default {
    name: "Counter",
    setup(){
        const state = reactive({
            count: 10,
            double: computed(()=>state.count*2)
        });

        const increment = () => {
            state.count = state.count * 2
        }


        return{
            state,
            increment
        }
    },

}
</script>

<style>

</style>