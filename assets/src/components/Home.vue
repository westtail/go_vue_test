<template>
    <div class="home">
        <h1>ホームページ</h1>
        <p>{{ name }}</p>
    </div>
</template>

<script>
export default {
    name: 'Home',
    props: {
        name: String
    }
}
</script>