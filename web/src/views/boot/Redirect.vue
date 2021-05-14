<template>
  <div></div>
</template>

<script>
export default {
  name: 'Redirect',
  components: {
  },
  data () {
    return {
    }
  },
  created () {
  },
  methods: {
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$router.replace(from.path)
    })
  }
}
</script>
