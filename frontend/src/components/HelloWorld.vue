<script setup>
import { reactive, ref } from "vue";
import axios from "axios";

// defineProps({
// 	msg: String,
// });

const data = reactive({ value: null });

axios.get("https://random-data-api.com/api/commerce/random_commerce").then((res) => {
	data.value = JSON.stringify(res.data, null, 2);
	console.log(data.value);
});
// .then((json) => (data.value = json));

console.log(data.value);

const state = reactive({ count: 0 });

const increment = () => {
	state.count++;
};
</script>

<template>
	<div class="text-center">
		<h1 class="text-blue-900 pt-6 text-2xl">{{ msg }}</h1>
		<p>count = {{ state.count }}</p>
		<pre>{{ data.value }}</pre>
		<button @click="increment" class="bg-slate-200">This will increment.</button>
	</div>
</template>

<!-- <style scoped></style> -->
