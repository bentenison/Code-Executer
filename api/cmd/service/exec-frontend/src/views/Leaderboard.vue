<template>
  <div
    class="flex flex-column justify-content-center"
    style="height: 100vh; margin-top: 14rem"
  >
    <div class="p-6 flex-1 w-100">
      <div class="flex align-items-center justify-content-between mb-6">
        <div class="text-2xl font-bold">EPIC.</div>
        <div class="flex align-items-center gap-4">
          <InputText
            class="px-4 py-2 rounded-lg"
            placeholder="Search"
            type="text"
          />
          <i class="fas fa-bell text-xl"> </i>
          <div class="flex align-items-center gap-2">
            <img
              alt="User profile picture"
              class="border-circle"
              height="40"
              src="https://storage.googleapis.com/a1aa/image/uBebxfZTHmhtRUbDzzrmos3QlJP2G2tbFkFcmDQppwKuA0zTA.jpg"
              width="40"
            />
            <span> Alex G. </span>
          </div>
        </div>
      </div>
      <div class="grid gap-2 mb-6 justify-content-center">
        <div class="col-3">
          <Card
            class="p-4 border-round-lg shadow-3"
            style="background: var(--p-green-100); color: var(--p-gray-600)"
          >
            <template #content>
              <div class="text-gray-600">Joined Members</div>
              <div class="text-2xl font-bold">346+</div>
            </template>
          </Card>
        </div>
        <div class="col-3">
          <Card
            class="p-4 border-round-lg shadow-3"
            style="background: var(--p-blue-100); color: var(--p-gray-600)"
          >
            <template #content>
              <div class="text-gray-600">Achieved goals</div>
              <div class="text-2xl font-bold">732+</div>
            </template>
          </Card>
        </div>
        <div class="col-3">
          <Card
            class="px-4 py-3 border-round-lg shadow-3"
            style="background: var(--p-yellow-100); color: var(--p-gray-600)"
          >
            <template #content>
              <div class="text-gray-600">
                Remaining time to completion
                <i class="pi pi-fire text-red-500"> </i>
              </div>
              <div class="text-2xl font-bold flex items-center gap-2">
                <span> 12 </span>
                <span> : </span>
                <span> 6 </span>
                <span> : </span>
                <span> 42 </span>
              </div>
              <div class="text-sm text-gray-500">
                Only the first three positions will be awarded prizes
              </div>
            </template>
          </Card>
        </div>
      </div>
      <div class="mb-6">
        <div class="text-xl font-bold mb-4">Current Leaders</div>
        <div
          class="flex align-items-center justify-content-center flex-wrap gap-4"
        >
          <Card class="p-4 border-round-lg shadow col-3">
            <template #content>
              <div class="flex align-items-center gap-4 mb-4">
                <img
                  alt="Theresa Webb profile picture"
                  class="border-circle"
                  height="50"
                  src="https://storage.googleapis.com/a1aa/image/jgJONbbtKcbLEtAD5MzFsW5GfkeHXJFj9AS7fg7IZH5iBonnA.jpg"
                  width="50"
                />
                <div>
                  <div class="font-bold">Theresa Webb</div>
                  <div class="">5.67532 pts.</div>
                </div>
                <i
                  class="pi pi-trophy"
                  style="font-size: 2rem; color: var(--p-yellow-500)"
                >
                </i>
              </div>
              <div
                class="flex align-items-center justify-content-between text-center"
              >
                <div>
                  <div class="font-bold">43</div>
                  <div class="text-gray-600">WINS</div>
                </div>
                <div>
                  <div class="font-bold">167</div>
                  <div class="text-gray-600">TASKS</div>
                </div>
                <div>
                  <div class="font-bold">476</div>
                  <div class="text-gray-600">ACH.</div>
                </div>
              </div>
            </template>
          </Card>
          <Card class="p-4 rounded-lg shadow col-3">
            <template #content>
              <div class="flex align-items-center gap-4 mb-4">
                <img
                  alt="Floyd Miles profile picture"
                  class="border-circle"
                  height="50"
                  src="https://storage.googleapis.com/a1aa/image/lDzXZGJB8hIAIxZ4aXagC8gjCpjWX8cnnOUveUKWSCkXA65JA.jpg"
                  width="50"
                />
                <div>
                  <div class="font-bold">Floyd Miles</div>
                  <div class="text-gray-600">5.44352 pts.</div>
                </div>
                <i
                  class="pi pi-trophy"
                  style="font-size: 2rem; color: var(--p-gray-400)"
                >
                </i>
              </div>
              <div
                class="flex align-items-center justify-content-between text-center"
              >
                <div>
                  <div class="font-bold">37</div>
                  <div class="text-gray-600">WINS</div>
                </div>
                <div>
                  <div class="font-bold">132</div>
                  <div class="text-gray-600">TASKS</div>
                </div>
                <div>
                  <div class="font-bold">482</div>
                  <div class="text-gray-600">ACH.</div>
                </div>
              </div>
            </template>
          </Card>
          <Card class="p-4 border-round-lg shadow col-3">
            <template #content>
              <div class="flex align-items-center gap-4 mb-4">
                <img
                  alt="Jacob Jones profile picture"
                  class="border-circle"
                  height="50"
                  width="50"
                  src="https://storage.googleapis.com/a1aa/image/EFcHomOYBE5WE58ZXCYnA8wc3E0snwoxipFgGBH9e1HYA65JA.jpg"
                />
                <div>
                  <div class="font-bold">Jacob Jones</div>
                  <div class="text-gray-600">4.36257 pts.</div>
                </div>
                <i
                  class="pi pi-trophy"
                  style="font-size: 2rem; color: var(--p-yellow-600)"
                >
                </i>
              </div>
              <div
                class="flex align-items-center justify-content-between text-center"
              >
                <div>
                  <div class="font-bold">32</div>
                  <div class="text-gray-600">WINS</div>
                </div>
                <div>
                  <div class="font-bold">68</div>
                  <div class="text-gray-600">TASKS</div>
                </div>
                <div>
                  <div class="font-bold">268</div>
                  <div class="text-gray-600">ACH.</div>
                </div>
              </div>
            </template>
          </Card>
        </div>
      </div>
    </div>
    <div
      class="card flex justify-content-center align-items-center border-round-lg"
    >
      <DataTable
        v-model:filters="filters"
        :value="customers"
        paginator
        :rows="10"
        dataKey="id"
        filterDisplay="row"
        style="width: 75%"
        :loading="loading"
        :globalFilterFields="[
          'name',
          'country.name',
          'representative.name',
          'status',
        ]"
      >
        <template #header>
          <div class="flex justify-end">
            <IconField>
              <InputIcon>
                <i class="pi pi-search mr-2" />
              </InputIcon>
              <InputText
                v-model="filters['global'].value"
                placeholder="Keyword Search"
              />
            </IconField>
          </div>
        </template>
        <template #empty> No customers found. </template>
        <template #loading> Loading customers data. Please wait. </template>
        <Column field="name" header="Name" style="min-width: 12rem">
          <template #body="{ data }">
            <div class="flex items-center gap-3">
              <img
                :alt="data.representative.name"
                :src="`https://primefaces.org/cdn/primevue/images/avatar/${data.representative.image}`"
                style="width: 36px"
              />
              <div class="flex flex-column gap-2">
                <p class="m-0 p-0">{{ data.name }}</p>
                <p class="m-0 p-0">#{{ data.id }}</p>
              </div>
            </div>
          </template>
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              placeholder="Search by name"
            />
          </template>
        </Column>
        <Column
          header="Country"
          filterField="country.name"
          style="min-width: 12rem"
        >
          <template #body="{ data }">
            <div class="flex items-center gap-2">
              <img
                alt="flag"
                src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
                :class="`flag flag-${data.country.code}`"
                style="width: 24px"
              />
              <span>{{ data.country.name }}</span>
            </div>
          </template>
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              placeholder="Search by country"
            />
          </template>
        </Column>
        <Column
          header="Agent"
          filterField="representative"
          :showFilterMenu="false"
          style="min-width: 14rem"
        >
          <template #body="{ data }">
            <div class="flex items-center gap-2">
              <img
                :alt="data.representative.name"
                :src="`https://primefaces.org/cdn/primevue/images/avatar/${data.representative.image}`"
                style="width: 32px"
              />
              <span>{{ data.representative.name }}</span>
            </div>
          </template>
          <template #filter="{ filterModel, filterCallback }">
            <MultiSelect
              v-model="filterModel.value"
              @change="filterCallback()"
              :options="representatives"
              optionLabel="name"
              placeholder="Any"
              style="min-width: 14rem"
              :maxSelectedLabels="1"
            >
              <template #option="slotProps">
                <div class="flex items-center gap-2">
                  <img
                    :alt="slotProps.option.name"
                    :src="`https://primefaces.org/cdn/primevue/images/avatar/${slotProps.option.image}`"
                    style="width: 32px"
                  />
                  <span>{{ slotProps.option.name }}</span>
                </div>
              </template>
            </MultiSelect>
          </template>
        </Column>
        <Column
          field="status"
          header="Status"
          :showFilterMenu="false"
          style="min-width: 12rem"
        >
          <template #body="{ data }">
            <Tag :value="data.status" :severity="getSeverity(data.status)" />
          </template>
          <template #filter="{ filterModel, filterCallback }">
            <Select
              v-model="filterModel.value"
              @change="filterCallback()"
              :options="statuses"
              placeholder="Select One"
              style="min-width: 12rem"
              :showClear="true"
            >
              <template #option="slotProps">
                <Tag
                  :value="slotProps.option"
                  :severity="getSeverity(slotProps.option)"
                />
              </template>
            </Select>
          </template>
        </Column>
        <Column
          field="verified"
          header="Verified"
          dataType="boolean"
          style="min-width: 6rem"
        >
          <template #body="{ data }">
            <i
              class="pi"
              :class="{
                'pi-check-circle text-green-500': data.verified,
                'pi-times-circle text-red-400': !data.verified,
              }"
            ></i>
          </template>
          <template #filter="{ filterModel, filterCallback }">
            <Checkbox
              v-model="filterModel.value"
              :indeterminate="filterModel.value === null"
              binary
              @change="filterCallback()"
            />
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
</template>

<script>
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import ColumnGroup from "primevue/columngroup"; // optional
import Row from "primevue/row"; // optional
import { FilterMatchMode } from "@primevue/core/api";
import MultiSelect from "primevue/multiselect";
import Select from "primevue/select";
export default {
  components: {
    DataTable,
    Column,
    ColumnGroup,
    Row,
    MultiSelect,
    Select,
  },
  data() {
    return {
      //   customers: null,
      filters: {
        global: { value: null, matchMode: FilterMatchMode.CONTAINS },
        name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
        "country.name": { value: null, matchMode: FilterMatchMode.STARTS_WITH },
        representative: { value: null, matchMode: FilterMatchMode.IN },
        status: { value: null, matchMode: FilterMatchMode.EQUALS },
        verified: { value: null, matchMode: FilterMatchMode.EQUALS },
      },
      representatives: [
        { name: "Amy Elsner", image: "amyelsner.png" },
        { name: "Anna Fali", image: "annafali.png" },
        { name: "Asiya Javayant", image: "asiyajavayant.png" },
        { name: "Bernardo Dominic", image: "bernardodominic.png" },
        { name: "Elwin Sharvill", image: "elwinsharvill.png" },
        { name: "Ioni Bowcher", image: "ionibowcher.png" },
        { name: "Ivan Magalhaes", image: "ivanmagalhaes.png" },
        { name: "Onyama Limba", image: "onyamalimba.png" },
        { name: "Stephen Shaw", image: "stephenshaw.png" },
        { name: "XuXue Feng", image: "xuxuefeng.png" },
      ],
      statuses: [
        "unqualified",
        "qualified",
        "new",
        "negotiation",
        "renewal",
        "proposal",
      ],
      loading: false,
      customers: [
        {
          id: 1000,
          name: "James Butt",
          country: {
            name: "Algeria",
            code: "dz",
          },
          company: "Benton, John B Jr",
          date: "2015-09-13",
          status: "unqualified",
          verified: true,
          activity: 17,
          representative: {
            name: "Ioni Bowcher",
            image: "ionibowcher.png",
          },
          balance: 70663,
        },
        {
          id: 1001,
          name: "Josephine Darakjy",
          country: {
            name: "Egypt",
            code: "eg",
          },
          company: "Chanay, Jeffrey A Esq",
          date: "2019-02-09",
          status: "negotiation",
          verified: true,
          activity: 0,
          representative: {
            name: "Amy Elsner",
            image: "amyelsner.png",
          },
          balance: 82429,
        },
        {
          id: 1002,
          name: "Art Venere",
          country: {
            name: "Panama",
            code: "pa",
          },
          company: "Chemel, James L Cpa",
          date: "2017-05-13",
          status: "qualified",
          verified: false,
          activity: 63,
          representative: {
            name: "Asiya Javayant",
            image: "asiyajavayant.png",
          },
          balance: 28334,
        },
        {
          id: 1003,
          name: "Lenna Paprocki",
          country: {
            name: "Slovenia",
            code: "si",
          },
          company: "Feltz Printing Service",
          date: "2020-09-15",
          status: "new",
          verified: false,
          activity: 37,
          representative: {
            name: "Xuxue Feng",
            image: "xuxuefeng.png",
          },
          balance: 88521,
        },
        {
          id: 1004,
          name: "Donette Foller",
          country: {
            name: "South Africa",
            code: "za",
          },
          company: "Printing Dimensions",
          date: "2016-05-20",
          status: "negotiation",
          verified: true,
          activity: 33,
          representative: {
            name: "Asiya Javayant",
            image: "asiyajavayant.png",
          },
          balance: 93905,
        },
        {
          id: 1005,
          name: "Simona Morasca",
          country: {
            name: "Egypt",
            code: "eg",
          },
          company: "Chapman, Ross E Esq",
          date: "2018-02-16",
          status: "qualified",
          verified: false,
          activity: 68,
          representative: {
            name: "Ivan Magalhaes",
            image: "ivanmagalhaes.png",
          },
          balance: 50041,
        },
        {
          id: 1006,
          name: "Mitsue Tollner",
          country: {
            name: "Paraguay",
            code: "py",
          },
          company: "Morlong Associates",
          date: "2018-02-19",
          status: "renewal",
          verified: true,
          activity: 54,
          representative: {
            name: "Ivan Magalhaes",
            image: "ivanmagalhaes.png",
          },
          balance: 58706,
        },
        {
          id: 1007,
          name: "Leota Dilliard",
          country: {
            name: "Serbia",
            code: "rs",
          },
          company: "Commercial Press",
          date: "2019-08-13",
          status: "renewal",
          verified: true,
          activity: 69,
          representative: {
            name: "Onyama Limba",
            image: "onyamalimba.png",
          },
          balance: 26640,
        },
      ],
    };
  },
  mounted() {
    // CustomerService.getCustomersMedium().then((data) => {
    //   this.customers = this.getCustomers(data);
    //   this.loading = false;
    // });
  },
  methods: {
    getCustomers(data) {
      return [...(data || [])].map((d) => {
        d.date = new Date(d.date);
        return d;
      });
    },
    getSeverity(status) {
      switch (status) {
        case "unqualified":
          return "danger";

        case "qualified":
          return "success";

        case "new":
          return "info";

        case "negotiation":
          return "warn";

        case "renewal":
          return null;
      }
    },
  },
};
</script>

<style lang="scss" scoped></style>
