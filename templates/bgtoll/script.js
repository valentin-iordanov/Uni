const Program = {
  delimiters: ['{%', '%}'],
  data() {
    return {
      carNum: "",
      info: {},
    }
  },
  methods: {
    incriment() {
      this.count = this.count + 1
    },
    async Info() {

      this.info = {}

      if (this.carNum == "") {
        return
      }
      const response = await fetch("https://check.bgtoll.bg/check/vignette/plate/BG/" + this.carNum)
      this.info = await response.json()

      console.log(this.info)

      

    }
  }
}


Vue.createApp(Program).mount('#Program')