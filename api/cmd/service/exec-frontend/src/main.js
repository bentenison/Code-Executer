// import 'primevue/resources/primevue.min.css';
import "primeflex/primeflex.css";
import "primeicons/primeicons.css";
import "./assets/theme/_dark.scss"
import "./assets/theme/layout.scss"
// import 'prismjs/themes/prism-coy.css';
// import './assets/styles/layout.scss';
// import './assets/demo/flags/flags.css';
import Aura from "@primevue/themes/aura";
import App from "./App.vue";
import mitt from "mitt";
import router from "./router/router"
import PrimeVue from "primevue/config";
import { createApp, reactive } from "vue";
import AutoComplete from "primevue/autocomplete";
import Accordion from "primevue/accordion";
import AccordionTab from "primevue/accordiontab";
import Avatar from "primevue/avatar";
import AvatarGroup from "primevue/avatargroup";
import Badge from "primevue/badge";
import BadgeDirective from "primevue/badgedirective";
import Button from "primevue/button";
import Breadcrumb from "primevue/breadcrumb";
import Calendar from "primevue/calendar";
import Card from "primevue/card";
import Carousel from "primevue/carousel";
// import Chart from 'primevue/chart';
import Checkbox from "primevue/checkbox";
import Chip from "primevue/chip";
// import Chips from 'primevue/chips';
import ColorPicker from "primevue/colorpicker";
import Column from "primevue/column";
import ConfirmDialog from "primevue/confirmdialog";
import ConfirmPopup from "primevue/confirmpopup";
import ConfirmationService from "primevue/confirmationservice";
import ContextMenu from "primevue/contextmenu";
import DataTable from "primevue/datatable";
import DataView from "primevue/dataview";
// import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions';
import Dialog from "primevue/dialog";
import Divider from "primevue/divider";
import Dropdown from "primevue/dropdown";
import Fieldset from "primevue/fieldset";
import FileUpload from "primevue/fileupload";
import Image from "primevue/image";
import InlineMessage from "primevue/inlinemessage";
import Inplace from "primevue/inplace";
import InputMask from "primevue/inputmask";
import InputNumber from "primevue/inputnumber";
import InputSwitch from "primevue/inputswitch";
import InputText from "primevue/inputtext";
import Knob from "primevue/knob";
import Galleria from "primevue/galleria";
import Listbox from "primevue/listbox";
import MegaMenu from "primevue/megamenu";
import Menu from "primevue/menu";
import Menubar from "primevue/menubar";
import Message from "primevue/message";
import MultiSelect from "primevue/multiselect";
import OrderList from "primevue/orderlist";
// import OrganizationChart from 'primevue/organizationchart';
import OverlayPanel from "primevue/overlaypanel";
import Paginator from "primevue/paginator";
import Panel from "primevue/panel";
import PanelMenu from "primevue/panelmenu";
import Password from "primevue/password";
import PickList from "primevue/picklist";
import ProgressBar from "primevue/progressbar";
import Rating from "primevue/rating";
import RadioButton from "primevue/radiobutton";
import Ripple from "primevue/ripple";
import SelectButton from "primevue/selectbutton";
import ScrollPanel from "primevue/scrollpanel";
import ScrollTop from "primevue/scrolltop";
import Slider from "primevue/slider";
import Sidebar from "primevue/sidebar";
import Skeleton from "primevue/skeleton";
import SplitButton from "primevue/splitbutton";
import Splitter from "primevue/splitter";
import SplitterPanel from "primevue/splitterpanel";
import Steps from "primevue/steps";
import StyleClass from "primevue/styleclass";
import TabMenu from "primevue/tabmenu";
import Tag from "primevue/tag";
import TieredMenu from "primevue/tieredmenu";
import Textarea from "primevue/textarea";
import Timeline from "primevue/timeline";
import Toast from "primevue/toast";
import ToastService from "primevue/toastservice";
import Toolbar from "primevue/toolbar";
import TabView from "primevue/tabview";
import TabPanel from "primevue/tabpanel";
import Tooltip from "primevue/tooltip";
import ToggleButton from "primevue/togglebutton";
import Tree from "primevue/tree";
import TreeSelect from "primevue/treeselect";
import TreeTable from "primevue/treetable";
import Popover from "primevue/popover";
// import TriStateCheckbox from 'primevue/tristatecheckbox';

// import CodeHighlight from './AppCodeHighlight';
// import BlockViewer from './BlockViewer';

const isProduction = process.env.NODE_ENV === "production";

// const options = {
//   isEnabled: true,
//   logLevel: isProduction ? 'error' : 'debug',
//   stringifyArguments: false,
//   showLogLevel: true,
//   showMethodName: true,
//   separator: '|',
//   showConsoleColors: true,
// }

const vm = createApp(App);
// const pinia = createPinia()
// pinia.use(piniaPluginPersistedstate)
vm.$router = router
// vm.$pinia = pinia
// vm.$confirm = ConfirmationService
vm.component("ConfirmDialog", ConfirmDialog);
// window.app = vm
// vm.config.globalProperties.$appState = reactive({
//   theme: "lara",
//   darkTheme: true,
// });

// vm.use(primeVue, { ripple: true, inputStyle: 'outlined' });
// vm.use(ConfirmationService);
vm.use(ToastService);
vm.use(router);

vm.directive("tooltip", Tooltip);
vm.directive("ripple", Ripple);
// vm.directive('code', CodeHighlight);
vm.directive("badge", BadgeDirective);
vm.directive("styleclass", StyleClass);

vm.component("Accordion", Accordion);
vm.component("AccordionTab", AccordionTab);
vm.component("AutoComplete", AutoComplete);
vm.component("Avatar", Avatar);
vm.component("AvatarGroup", AvatarGroup);
vm.component("Badge", Badge);
vm.component("Breadcrumb", Breadcrumb);
vm.component("Button", Button);
vm.component("Calendar", Calendar);
vm.component("Card", Card);
vm.component("Carousel", Carousel);
// vm.component('Chart', Chart);
vm.component("Checkbox", Checkbox);
vm.component("Chip", Chip);
vm.component('Popover',Popover)
// vm.component('Chips', Chips);
vm.component("ColorPicker", ColorPicker);
vm.component("Column", Column);
// vm.component('ConfirmDialog', ConfirmDialog);
vm.component("ConfirmPopup", ConfirmPopup);
vm.component("ContextMenu", ContextMenu);
vm.component("DataTable", DataTable);
vm.component("DataView", DataView);
// vm.component('DataViewLayoutOptions', DataViewLayoutOptions);
vm.component("Dialog", Dialog);
vm.component("Divider", Divider);
vm.component("Dropdown", Dropdown);
vm.component("Fieldset", Fieldset);
vm.component("FileUpload", FileUpload);
vm.component("Image", Image);
vm.component("InlineMessage", InlineMessage);
vm.component("Inplace", Inplace);
vm.component("InputMask", InputMask);
vm.component("InputNumber", InputNumber);
vm.component("InputSwitch", InputSwitch);
vm.component("InputText", InputText);
vm.component("Galleria", Galleria);
vm.component("Knob", Knob);
vm.component("Listbox", Listbox);
vm.component("MegaMenu", MegaMenu);
vm.component("Menu", Menu);
vm.component("Menubar", Menubar);
vm.component("Message", Message);
vm.component("MultiSelect", MultiSelect);
vm.component("OrderList", OrderList);
// vm.component('OrganizationChart', OrganizationChart);
vm.component("OverlayPanel", OverlayPanel);
vm.component("Paginator", Paginator);
vm.component("Panel", Panel);
vm.component("PanelMenu", PanelMenu);
vm.component("Password", Password);
vm.component("PickList", PickList);
vm.component("ProgressBar", ProgressBar);
vm.component("RadioButton", RadioButton);
vm.component("Rating", Rating);
vm.component("SelectButton", SelectButton);
vm.component("ScrollPanel", ScrollPanel);
vm.component("ScrollTop", ScrollTop);
vm.component("Slider", Slider);
vm.component("Sidebar", Sidebar);
vm.component("Skeleton", Skeleton);
vm.component("SplitButton", SplitButton);
vm.component("Splitter", Splitter);
vm.component("SplitterPanel", SplitterPanel);
vm.component("Steps", Steps);
vm.component("TabMenu", TabMenu);
vm.component("TabView", TabView);
vm.component("TabPanel", TabPanel);
vm.component("Tag", Tag);
vm.component("Textarea", Textarea);
vm.component("TieredMenu", TieredMenu);
vm.component("Timeline", Timeline);
vm.component("Toast", Toast);
vm.component("Toolbar", Toolbar);
vm.component("ToggleButton", ToggleButton);
vm.component("Tree", Tree);
vm.component("TreeSelect", TreeSelect);
vm.component("TreeTable", TreeTable);
// vm.component('TriStateCheckbox', TriStateCheckbox);

// vm.component('BlockViewer', BlockViewer);
// vm.use(PiniaVuePlugin)
// vm.use(pinia)
// vm.use(i18n)
// vm.use(primeVue)
// vm.use(router)
vm.use(ConfirmationService);
vm.use(ConfirmDialog);

// import Aura from "@primevue/themes/lara";
const emitter = mitt();

vm.config.globalProperties.emitter = emitter;
vm.provide('emitter', emitter);  
vm.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: ".app-dark",
    },
  },
});
vm.mount("#app");
