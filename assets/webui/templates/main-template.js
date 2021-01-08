// Note that these are backticks.

const MainTemplate = String.raw`
  <q-layout view="hHh lpR fFf">
  
    <q-header bordered class="bg-primary text-white">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="left = !left" />
        <q-toolbar-title>
          {{titlePrefix}}
          <q-avatar>
            <img src="icons/go-logo-blue.svg">
          </q-avatar>
          with
          <q-avatar square>
            <img src="icons/vue-logo.png">
          </q-avatar>
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
  
    <q-drawer show-if-above v-model="left" side="left" bordered>
      <q-list>
        <q-item>
          <q-input standout bottom-slots v-model="name" label="Name" @keyup.enter="setName">
            <template v-slot:before>
              <q-icon name="account_circle" />
            </template>
            <template v-slot:append>
              <q-btn round dense flat icon="check" @click="setName" />
            </template>
          </q-input>
        </q-item>
        <q-item>
          <q-btn label="check version" color="primary" @click="checkVersion" class="full-width"></q-btn>
        </q-item>
      </q-list>
    </q-drawer>
  
    <q-page-container>
      <div class="q-pa-md row justify-center">
        <div style="width: 100%">
          <q-chat-message v-if="message">{{message}}</q-chat-message>
        </div>
      </div>
    </q-page-container>
  
  </q-layout>
`

export { MainTemplate }