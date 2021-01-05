import { MainTemplate } from './templates/main-template.js'
import { InitAPI } from './api.js'
const queryString = window.location.search
const urlParams = new URLSearchParams(queryString)
var apiURL = "../api/"
if (urlParams.has('api')) {
  apiURL = urlParams.get('api')
}
const API = InitAPI(apiURL)

var app = new Vue({
  el: '#q-app',
  data: {
    titlePrefix: "",
    name: "",
    message: "",
    left: false,
  },
  methods: {
    checkVersion: function () {
      API.getVersion().
        then(function (result) {
          app.$q.notify('Running on v' + result)
        }).catch(function (error) {
          app.$q.notify('Looks like there was an API problem: ' + error)
        })
    },
    setName: function () {
      console.log(app.name)
      API.setName(app.name).
        then(function (result) {
          app.message = result
        }).catch(function (error) {
          app.$q.notify('Looks like there was an API problem: ' + error)
        })
    },
  },
  template: MainTemplate
})

API.getTitlePrefix().
  then(function (result) {
    app.titlePrefix = result
  }).catch(function (error) {
    app.$q.notify('Looks like there was an API problem: ' + error)
  })
