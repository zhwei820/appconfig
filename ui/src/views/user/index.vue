<template>
  <div class="app-container" >

    <div class="ui grid">
      <div class="ui left aligned nine wide column">
        <div class="ui labeled icon input">
          <div class="ui label">Search:</div>
          <input v-model="searchFor" class="ui input" @keyup.enter="setFilter">
          <i class="search icon"/>
        </div>
        <button class="ui button primary" @click="setFilter">Go</button>
        <button class="ui button" @click="resetFilter">Reset</button>
      </div>
    </div>

    <div :class="[{'vuetable-wrapper ui basic segment': true}, loading]">
      <vuetable
        ref="vuetable"
        :fields="fields"
        :table-height="tableHeight"
        :max-table-height="maxTableHeight"
        :row-class="rowClassCB"
        :sort-order="sortOrder"
        :multi-sort="multiSort"
        :api-mode="false"
        :data="transformed"
        track-by="id"
        pagination-path="pagination"
        @vuetable:pagination-data="onPaginationData"
        @vuetable:load-success="onLoadSuccess"
        @vuetable:loading="showLoader"
        @vuetable:loaded="hideLoader"
        @vuetable:cell-clicked="onCellClicked"
        @vuetable:initialized="onInitialized"
        @vuetable:orderby="orderby"
        @vuetable:data-reset="onDataReset">
        <template slot="actions" slot-scope="props">
          <div class="custom-actions">
            <button class="ui compact teal button" @click="onAction('view-item', props.rowData, props.rowIndex)">
              <i class="zoom icon"/>
            </button>
            <button class="ui compact blue button" @click="onAction('edit-item', props.rowData, props.rowIndex)">
              <i class="edit icon"/>
            </button>
            <button class="ui compact red button" @click="onAction('delete-item', props.rowData, props.rowIndex)">
              <i class="delete icon"/>
            </button>
          </div>
        </template>
      </vuetable>
      <div class="vuetable-pagination ui bottom attached segment grid">
        <vuetable-pagination-info ref="paginationInfo" :info-template="paginationInfoTemplate"/>
        <component ref="pagination" :is="paginationComponent" @vuetable-pagination:change-page="onChangePage"/>
      </div>
    </div>
  </div>
</template>

<script>
import Vuetable from 'vuetable-2/src/components/Vuetable'
import VuetablePagination from 'vuetable-2/src/components/VuetablePagination'
import VuetablePaginationInfo from 'vuetable-2/src/components/VuetablePaginationInfo'
// import { fetchList, createUser, updateUser } from '@/api/user'
import { fetchList } from '@/api/user'
import { preg_quote, getSortObj } from '@/utils/index'
import { Message } from 'element-ui'

const lang = {
  'create_at': 'create_at'
}

const tableColumns = [
  {
    name: '__checkbox',
    width: '30px',
    title: 'checkbox',
    titleClass: 'center aligned',
    dataClass: 'center aligned'
  },
  {
    name: 'id',
    title: '<i class="unordered list icon"></i> Id',
    dataClass: 'center aligned',
    width: '100px'
  },
  {
    name: 'username',
    title: '<i class="book icon"></i> username',
    sortField: 'username',
    width: '150px'
  },
  {
    name: 'group',
    title: '<i class="mail outline icon"></i> group',
    width: '200px',
    visible: true,
    titleClass: 'center aligned',
    dataClass: 'center aligned',
    callback: 'group'
  },
  {
    name: 'create_at',
    title: (nameOnly = false) => {
      return nameOnly
        ? lang['create_at']
        : `<i class="orange birthday icon"></i> ${lang['create_at']}`
    },
    sortField: 'create_at',
    width: '150px',
    callback: 'formatDate|D/MM/Y'
  },
  {
    name: 'gender',
    title: 'Gender',
    titleClass: 'center aligned',
    dataClass: 'center aligned',
    callback: 'gender',
    width: '100px'
  },
  {
    name: '__slot:actions', // <----
    title: 'Actions',
    titleClass: 'center aligned',
    dataClass: 'center aligned',
    width: '150px'
  }
]

export default {
  name: 'ComplexTable',
  components: {
    Vuetable,
    VuetablePagination,
    VuetablePaginationInfo
  },
  data: function() {
    return {
      transformed: {},
      selectedTo: [],
      loading: '',
      searchFor: '',
      moreParams: { aa: 1111, bb: 222 },
      fields: tableColumns,
      tableHeight: 'auto',
      maxTableHeight: '700px',
      vuetableFields: false,
      sortOrder: [],
      multiSort: true,
      paginationComponent: 'vuetable-pagination',
      perPage: 5,
      page: 1,
      totalPage: 1,
      paginationInfoTemplate: '',
      // paginationInfoTemplate: 'Showing record: {from} to {to} from {total} item(s)',
      lang: lang
    }
  },
  watch: {
    'paginationComponent'(val, oldVal) {
      this.$nextTick(function() {
        this.$refs.pagination.setPaginationData(this.$refs.vuetable.tablePagination)
      })
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    updateParams() {
      return { ...this.moreParams, page: this.page, per_page: this.perPage }
    },
    fetchData() {
      this.showLoader()
      fetchList({ ...this.updateParams(), ...getSortObj(this.sortOrder) }).then(response => {
        this.transformData(response)
        this.hideLoader()
      })
    },
    transformData(data) {
      this.transformed = {}

      this.totalPage = Math.ceil(data.count / this.perPage)

      this.transformed.pagination = {
        total: data.count,
        per_page: this.perPage,
        current_page: this.page,
        last_page: this.totalPage
      }

      this.transformed.data = data.data
    },
    showLoader() {
      this.loading = 'loading'
    },
    hideLoader() {
      this.loading = ''
    },
    formatDate(value, fmt) {
      if (value === null) return ''
      fmt = (typeof (fmt) === 'undefined') ? 'D MMM YYYY' : fmt
      return value
    },
    gender(value) {
      return '<span class="ui pink label"><i class="female icon"></i></span> <span class="el-tag el-tag--primary el-tag--mini">新-新提交</span>'
    },
    group(value) {
      return value + '<span class="el-tag el-tag--danger el-tag--mini">2</span>'
    },
    setFilter() {
      this.moreParams['username__icontains'] = this.searchFor
      this.$nextTick(function() {
        this.fetchData()
      })
    },
    resetFilter() {
      this.searchFor = ''
      this.setFilter()
    },
    highlight(needle, haystack) {
      return haystack.replace(
        new RegExp('(' + preg_quote(needle) + ')', 'ig'),
        '<span class="highlight">$1</span>'
      )
    },
    rowClassCB(data, index) {
      return (index % 2) === 0 ? 'odd' : 'even'
    },
    orderby(field, event) {
      this.fetchData()
      console.log('orderby', this.sortOrder)
    },
    onCellClicked(data, field, event) {
      console.log('cellClicked', field.name, this.$refs.vuetable.selectedTo)
    },
    onCellDoubleClicked(data, field, event) {
      console.log('cellDoubleClicked:', field.name)
    },
    onCellRightClicked(data, field, event) {
      console.log('cellRightClicked:', field.name)
    },
    onLoadSuccess(response) {
      // set pagination data to pagination-info component
      this.$refs.paginationInfo.setPaginationData(response.data)

      const data = response.data.data
      if (this.searchFor !== '') {
        for (const n in data) {
          data[n].name = this.highlight(this.searchFor, data[n].name)
          data[n].email = this.highlight(this.searchFor, data[n].email)
        }
      }
    },
    onLoadError(response) {
      if (response.status === 400) {
        Message.warning('Something\'s Wrong!', response.data.message, 'error')
      } else {
        Message.warning('Oops', 'oooo', 'error')
      }
    },
    onPaginationData(tablePagination) {
      // debugger
      console.log(tablePagination)
      this.$refs.paginationInfo.setPaginationData(tablePagination)
      this.$refs.pagination.setPaginationData(tablePagination)
    },
    onChangePage(page) {
      if (page === 'next') {
        page = this.page + 1
      }
      if (page === 'prev') {
        page = this.page - 1
      }
      if (this.page === page || page <= 0 || page > this.totalPage) {
        return
      }
      this.page = page
      console.log('changePage', this.page)
      this.fetchData()
    },
    onInitialized(fields) {
      console.log('onInitialized', fields)
      this.vuetableFields = fields
    },
    onDataReset() {
      console.log('onDataReset')
      this.$refs.paginationInfo.resetData()
      this.$refs.pagination.resetData()
    },
    onAction(action, data, index) {
      console.log('slot) action: ' + action, data, data.name, index)
    },
    orderBy(field, event) {
      debugger
      console.log(field)
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  @import "src/styles/extra.scss";
</style>
