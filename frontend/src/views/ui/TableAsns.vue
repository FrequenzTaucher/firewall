<template>
  <div id="pageTable">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex sm12 v-if="errors.length > 0">
          <div class="v-alert error">
            <i
              aria-hidden="true"
              class="v-icon fa fa fa fa-exclamation-triangle theme--light v-alert__icon"
            ></i>
            <div>
              <ul>
                <li v-for="error in errors">
                  {{ error }}
                </li>
              </ul>
            </div>
          </div>
        </v-flex>
        <v-flex lg12>
          <v-card>
            <v-toolbar card color="white">
              <v-text-field
                flat
                solo
                prepend-icon="search"
                placeholder="Type something"
                v-model="search"
                hide-details
                class="hidden-sm-and-down"
              ></v-text-field>
              <v-btn icon>
                <v-icon>filter_list</v-icon>
              </v-btn>

              <v-dialog v-model="dialog" max-width="500px">
                <template v-slot:activator="{ on }">
                  <v-btn color="primary" dark class="mb-2" v-on="on">New Item</v-btn>
                </template>
                <v-card>
                  <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text>
                    <v-container grid-list-md>
                      <v-layout wrap>
                        <v-flex xs12 sm6 md4>
                          <v-text-field v-model="editedItem.value" label="Value"></v-text-field>
                        </v-flex>
                        <v-flex xs12 sm6 md4>
                          <v-select
                                  item-text="visible"
                                  item-value="value"
                                  v-model="editedItem.status"
                                  :items="statusOptions"
                                  label="Status"
                          ></v-select>
                        </v-flex>
                      </v-layout>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
                    <v-btn color="blue darken-1" flat @click="save">Save</v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
            <v-divider></v-divider>
            <v-card-text class="pa-0">
              <v-data-table
                :headers="headers"
                :search="search"
                :items="items"
                :rows-per-page-items="[10, 25, 50, { text: 'All', value: -1 }]"
                class="elevation-1"
                item-key="name"
              >
                <template slot="items" slot-scope="props">
                  <td>{{ props.item.value }}</td>
                  <td>
                    <span v-show="props.item.status ===  'deny'" tabindex="0" class="v-chip theme--light red white--text">
                      <span class="v-chip__content">
                      <div class="v-avatar" style="height: 48px; width: 48px;">
                        <i aria-hidden="true" class="v-icon material-icons theme--light white--text">block</i>
                      </div>
                      Deny
                    </span>
                    </span>
                    <span v-show="props.item.status ===  'allow'" tabindex="0" class="v-chip theme--light green white--text">
                      <span class="v-chip__content">
                      <div class="v-avatar" style="height: 48px; width: 48px;">
                        <i aria-hidden="true" class="v-icon material-icons theme--light white--text">check_circle</i>
                      </div>
                      Allow
                    </span>
                    </span>
                  </td>
                  <td>
                    <v-btn
                      depressed
                      outline
                      icon
                      fab
                      dark
                      color="primary"
                      small
                    >
                      <v-icon @click="editItem(props.item)">edit</v-icon>
                      <!-- {{ props.item._id }} -->
                    </v-btn>
                    <v-btn depressed outline icon fab dark color="pink" small>
                      <v-icon @click="deleteItem(props.item)">delete</v-icon>
                    </v-btn>
                  </td>
                </template>
              </v-data-table>
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import ApiClient from "../../api/client";

const client = new ApiClient();

export default {
  data: function () {
    return {
      dialog: false,
      search: "",
      headers: [
        {
          text: "Value",
          value: "value"
        },
        {
          text: "Status",
          value: "status"
        },
        {
          text: "Action",
          value: ""
        }
      ],
      editedIndex: -1,
      editedItem: {
        value: '',
        status: ''
      },
      defaultItem: {
        value: '',
        status: ''
      },
      items: [],
      errors: [],
      statusOptions: [
        {
          visible: "Allow",
          value: "allow"
        },
        {
          visible: "Deny",
          value: "deny"
        }
      ],
      crudUrls: {
        delete: this.getApiUrl(this.$route.path).replace('all', 'delete') + '/',
        create: this.getApiUrl(this.$route.path).replace('all', 'create'),
        update: this.getApiUrl(this.$route.path).replace('all', 'update'),
      }
    };
  },
  computed: {
    formTitle () {
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    }
  },

  watch: {
    dialog (val) {
      val || this.close()
    }
  },
  created() {
    client
      .all(this.getApiUrl(this.$route.path))
      .then(response => {
        response.status === 200
          ? (this.items = response.data)
          : (this.errors = response.data);
      })
      .catch(e => {
        this.errors.push(e);
      });
  },
  methods: {
    editItem (item) {
      this.editedIndex = this.items.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true
    },

    deleteItem (item) {
      const index = this.items.indexOf(item);
      confirm('Are you sure you want to delete this item?');

      client.delete(this.crudUrls.delete + item._id)
              .then(response => {
                response.status === 200
                        ? (this.items.splice(index, 1))
                        : (this.errors = response.data);
              })
              .catch(e => {
                this.errors.push(e);
              });

    },

    close () {
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1
      }, 300)
    },

    save () {

      if (this.editedIndex > -1) {
        let url = this.crudUrls.update;
        let payload = {_id: this.editedItem._id, value: this.editedItem.value, status: this.editedItem.status};

        client.update(url, payload)
                .then(response => {
                  response.status === 200
                          ? (Object.assign(this.items[this.editedIndex], this.editedItem))
                          : (this.errors = response.data);
                })
                .catch(e => {
                  this.errors.push(e);
                });
      } else {
        let url = this.crudUrls.create;
        let payload = {value: this.editedItem.value, status: this.editedItem.status};

        client.create(url, payload)
                .then(response => {
                  response.status === 200
                          ? (this.items.push(this.editedItem))
                          : (this.errors = response.data);
                })
                .catch(e => {
                  this.errors.push(e);
                });


      }

      this.close()
    }
  }
};
</script>
