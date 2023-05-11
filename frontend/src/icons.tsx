export {}
import { ICON_TYPES, ValueOf } from '@elastic/eui'

import { appendIconComponentCache } from '@elastic/eui/es/components/icon/icon'

import { icon as accessibility } from '@elastic/eui/es/components/icon/assets/accessibility'
import { icon as aggregate } from '@elastic/eui/es/components/icon/assets/aggregate'
import { icon as alert } from '@elastic/eui/es/components/icon/assets/alert'
import { icon as analyze_event } from '@elastic/eui/es/components/icon/assets/analyze_event'
import { icon as analyzeEvent } from '@elastic/eui/es/components/icon/assets/analyzeEvent'
import { icon as annotation } from '@elastic/eui/es/components/icon/assets/annotation'
import { icon as apm_trace } from '@elastic/eui/es/components/icon/assets/apm_trace'
import { icon as app_add_data } from '@elastic/eui/es/components/icon/assets/app_add_data'
import { icon as app_advanced_settings } from '@elastic/eui/es/components/icon/assets/app_advanced_settings'
import { icon as app_agent } from '@elastic/eui/es/components/icon/assets/app_agent'
import { icon as app_apm } from '@elastic/eui/es/components/icon/assets/app_apm'
import { icon as app_app_search } from '@elastic/eui/es/components/icon/assets/app_app_search'
import { icon as app_auditbeat } from '@elastic/eui/es/components/icon/assets/app_auditbeat'
import { icon as app_canvas } from '@elastic/eui/es/components/icon/assets/app_canvas'
import { icon as app_cases } from '@elastic/eui/es/components/icon/assets/app_cases'
import { icon as app_code } from '@elastic/eui/es/components/icon/assets/app_code'
import { icon as app_console } from '@elastic/eui/es/components/icon/assets/app_console'
import { icon as app_cross_cluster_replication } from '@elastic/eui/es/components/icon/assets/app_cross_cluster_replication'
import { icon as app_dashboard } from '@elastic/eui/es/components/icon/assets/app_dashboard'
import { icon as app_devtools } from '@elastic/eui/es/components/icon/assets/app_devtools'
import { icon as app_discover } from '@elastic/eui/es/components/icon/assets/app_discover'
import { icon as app_ems } from '@elastic/eui/es/components/icon/assets/app_ems'
import { icon as app_filebeat } from '@elastic/eui/es/components/icon/assets/app_filebeat'
import { icon as app_fleet } from '@elastic/eui/es/components/icon/assets/app_fleet'
import { icon as app_gis } from '@elastic/eui/es/components/icon/assets/app_gis'
import { icon as app_graph } from '@elastic/eui/es/components/icon/assets/app_graph'
import { icon as app_grok } from '@elastic/eui/es/components/icon/assets/app_grok'
import { icon as app_heartbeat } from '@elastic/eui/es/components/icon/assets/app_heartbeat'
import { icon as app_index_management } from '@elastic/eui/es/components/icon/assets/app_index_management'
import { icon as app_index_pattern } from '@elastic/eui/es/components/icon/assets/app_index_pattern'
import { icon as app_index_rollup } from '@elastic/eui/es/components/icon/assets/app_index_rollup'
import { icon as app_lens } from '@elastic/eui/es/components/icon/assets/app_lens'
import { icon as app_logs } from '@elastic/eui/es/components/icon/assets/app_logs'
import { icon as app_management } from '@elastic/eui/es/components/icon/assets/app_management'
import { icon as app_metricbeat } from '@elastic/eui/es/components/icon/assets/app_metricbeat'
import { icon as app_metrics } from '@elastic/eui/es/components/icon/assets/app_metrics'
import { icon as app_ml } from '@elastic/eui/es/components/icon/assets/app_ml'
import { icon as app_monitoring } from '@elastic/eui/es/components/icon/assets/app_monitoring'
import { icon as app_notebook } from '@elastic/eui/es/components/icon/assets/app_notebook'
import { icon as app_packetbeat } from '@elastic/eui/es/components/icon/assets/app_packetbeat'
import { icon as app_pipeline } from '@elastic/eui/es/components/icon/assets/app_pipeline'
import { icon as app_recently_viewed } from '@elastic/eui/es/components/icon/assets/app_recently_viewed'
import { icon as app_reporting } from '@elastic/eui/es/components/icon/assets/app_reporting'
import { icon as app_saved_objects } from '@elastic/eui/es/components/icon/assets/app_saved_objects'
import { icon as app_search_profiler } from '@elastic/eui/es/components/icon/assets/app_search_profiler'
// import { icon as app_security_analytics } from '@elastic/eui/es/components/icon/assets/app_security_analytics'
import { icon as app_security } from '@elastic/eui/es/components/icon/assets/app_security'
import { icon as apps } from '@elastic/eui/es/components/icon/assets/apps'
import { icon as app_spaces } from '@elastic/eui/es/components/icon/assets/app_spaces'
import { icon as app_sql } from '@elastic/eui/es/components/icon/assets/app_sql'
import { icon as app_timelion } from '@elastic/eui/es/components/icon/assets/app_timelion'
import { icon as app_upgrade_assistant } from '@elastic/eui/es/components/icon/assets/app_upgrade_assistant'
import { icon as app_uptime } from '@elastic/eui/es/components/icon/assets/app_uptime'
import { icon as app_users_roles } from '@elastic/eui/es/components/icon/assets/app_users_roles'
import { icon as app_visualize } from '@elastic/eui/es/components/icon/assets/app_visualize'
import { icon as app_watches } from '@elastic/eui/es/components/icon/assets/app_watches'
import { icon as app_workplace_search } from '@elastic/eui/es/components/icon/assets/app_workplace_search'
import { icon as arrow_down } from '@elastic/eui/es/components/icon/assets/arrow_down'
import { icon as arrowEnd } from '@elastic/eui/es/components/icon/assets/arrowEnd'
import { icon as arrow_left } from '@elastic/eui/es/components/icon/assets/arrow_left'
import { icon as arrow_right } from '@elastic/eui/es/components/icon/assets/arrow_right'
import { icon as arrowStart } from '@elastic/eui/es/components/icon/assets/arrowStart'
import { icon as arrow_up } from '@elastic/eui/es/components/icon/assets/arrow_up'
import { icon as asterisk } from '@elastic/eui/es/components/icon/assets/asterisk'
import { icon as beaker } from '@elastic/eui/es/components/icon/assets/beaker'
import { icon as bell } from '@elastic/eui/es/components/icon/assets/bell'
import { icon as bellSlash } from '@elastic/eui/es/components/icon/assets/bellSlash'
import { icon as bolt } from '@elastic/eui/es/components/icon/assets/bolt'
import { icon as boxes_horizontal } from '@elastic/eui/es/components/icon/assets/boxes_horizontal'
import { icon as boxes_vertical } from '@elastic/eui/es/components/icon/assets/boxes_vertical'
import { icon as branch } from '@elastic/eui/es/components/icon/assets/branch'
import { icon as branchUser } from '@elastic/eui/es/components/icon/assets/branchUser'
import { icon as broom } from '@elastic/eui/es/components/icon/assets/broom'
import { icon as brush } from '@elastic/eui/es/components/icon/assets/brush'
import { icon as bug } from '@elastic/eui/es/components/icon/assets/bug'
import { icon as bullseye } from '@elastic/eui/es/components/icon/assets/bullseye'
import { icon as calendar } from '@elastic/eui/es/components/icon/assets/calendar'
import { icon as checkInCircleFilled } from '@elastic/eui/es/components/icon/assets/checkInCircleFilled'
import { icon as check } from '@elastic/eui/es/components/icon/assets/check'
import { icon as cheer } from '@elastic/eui/es/components/icon/assets/cheer'
import { icon as clock } from '@elastic/eui/es/components/icon/assets/clock'
import { icon as cloudDrizzle } from '@elastic/eui/es/components/icon/assets/cloudDrizzle'
import { icon as cloudStormy } from '@elastic/eui/es/components/icon/assets/cloudStormy'
import { icon as cloudSunny } from '@elastic/eui/es/components/icon/assets/cloudSunny'
import { icon as color } from '@elastic/eui/es/components/icon/assets/color'
import { icon as compute } from '@elastic/eui/es/components/icon/assets/compute'
import { icon as console } from '@elastic/eui/es/components/icon/assets/console'
import { icon as continuityAboveBelow } from '@elastic/eui/es/components/icon/assets/continuityAboveBelow'
import { icon as continuityAbove } from '@elastic/eui/es/components/icon/assets/continuityAbove'
import { icon as continuityBelow } from '@elastic/eui/es/components/icon/assets/continuityBelow'
import { icon as continuityWithin } from '@elastic/eui/es/components/icon/assets/continuityWithin'
import { icon as controls_horizontal } from '@elastic/eui/es/components/icon/assets/controls_horizontal'
import { icon as controls_vertical } from '@elastic/eui/es/components/icon/assets/controls_vertical'
import { icon as copy_clipboard } from '@elastic/eui/es/components/icon/assets/copy_clipboard'
import { icon as copy } from '@elastic/eui/es/components/icon/assets/copy'
import { icon as crosshairs } from '@elastic/eui/es/components/icon/assets/crosshairs'
import { icon as cross } from '@elastic/eui/es/components/icon/assets/cross'
import { icon as currency } from '@elastic/eui/es/components/icon/assets/currency'
import { icon as cut } from '@elastic/eui/es/components/icon/assets/cut'
import { icon as database } from '@elastic/eui/es/components/icon/assets/database'
import { icon as desktop } from '@elastic/eui/es/components/icon/assets/desktop'
import { icon as documentation } from '@elastic/eui/es/components/icon/assets/documentation'
import { icon as documentEdit } from '@elastic/eui/es/components/icon/assets/documentEdit'
import { icon as document } from '@elastic/eui/es/components/icon/assets/document'
import { icon as documents } from '@elastic/eui/es/components/icon/assets/documents'
import { icon as dot } from '@elastic/eui/es/components/icon/assets/dot'
import { icon as doubleArrowLeft } from '@elastic/eui/es/components/icon/assets/doubleArrowLeft'
import { icon as doubleArrowRight } from '@elastic/eui/es/components/icon/assets/doubleArrowRight'
import { icon as download } from '@elastic/eui/es/components/icon/assets/download'
import { icon as editor_align_center } from '@elastic/eui/es/components/icon/assets/editor_align_center'
import { icon as editor_align_left } from '@elastic/eui/es/components/icon/assets/editor_align_left'
import { icon as editor_align_right } from '@elastic/eui/es/components/icon/assets/editor_align_right'
import { icon as editor_bold } from '@elastic/eui/es/components/icon/assets/editor_bold'
import { icon as editor_checklist } from '@elastic/eui/es/components/icon/assets/editor_checklist'
import { icon as editor_code_block } from '@elastic/eui/es/components/icon/assets/editor_code_block'
import { icon as editor_comment } from '@elastic/eui/es/components/icon/assets/editor_comment'
import { icon as editorDistributeHorizontal } from '@elastic/eui/es/components/icon/assets/editorDistributeHorizontal'
import { icon as editorDistributeVertical } from '@elastic/eui/es/components/icon/assets/editorDistributeVertical'
import { icon as editor_heading } from '@elastic/eui/es/components/icon/assets/editor_heading'
import { icon as editor_italic } from '@elastic/eui/es/components/icon/assets/editor_italic'
import { icon as editorItemAlignBottom } from '@elastic/eui/es/components/icon/assets/editorItemAlignBottom'
import { icon as editorItemAlignCenter } from '@elastic/eui/es/components/icon/assets/editorItemAlignCenter'
import { icon as editorItemAlignLeft } from '@elastic/eui/es/components/icon/assets/editorItemAlignLeft'
import { icon as editorItemAlignMiddle } from '@elastic/eui/es/components/icon/assets/editorItemAlignMiddle'
import { icon as editorItemAlignRight } from '@elastic/eui/es/components/icon/assets/editorItemAlignRight'
import { icon as editorItemAlignTop } from '@elastic/eui/es/components/icon/assets/editorItemAlignTop'
import { icon as editor_link } from '@elastic/eui/es/components/icon/assets/editor_link'
import { icon as editor_ordered_list } from '@elastic/eui/es/components/icon/assets/editor_ordered_list'
import { icon as editorPositionBottomLeft } from '@elastic/eui/es/components/icon/assets/editorPositionBottomLeft'
import { icon as editorPositionBottomRight } from '@elastic/eui/es/components/icon/assets/editorPositionBottomRight'
import { icon as editorPositionTopLeft } from '@elastic/eui/es/components/icon/assets/editorPositionTopLeft'
import { icon as editorPositionTopRight } from '@elastic/eui/es/components/icon/assets/editorPositionTopRight'
import { icon as editor_redo } from '@elastic/eui/es/components/icon/assets/editor_redo'
import { icon as editor_strike } from '@elastic/eui/es/components/icon/assets/editor_strike'
import { icon as editor_table } from '@elastic/eui/es/components/icon/assets/editor_table'
import { icon as editor_underline } from '@elastic/eui/es/components/icon/assets/editor_underline'
import { icon as editor_undo } from '@elastic/eui/es/components/icon/assets/editor_undo'
import { icon as editor_unordered_list } from '@elastic/eui/es/components/icon/assets/editor_unordered_list'
import { icon as email } from '@elastic/eui/es/components/icon/assets/email'
import { icon as empty } from '@elastic/eui/es/components/icon/assets/empty'
import { icon as eql } from '@elastic/eui/es/components/icon/assets/eql'
import { icon as eraser } from '@elastic/eui/es/components/icon/assets/eraser'
import { icon as exit } from '@elastic/eui/es/components/icon/assets/exit'
import { icon as expand } from '@elastic/eui/es/components/icon/assets/expand'
import { icon as expandMini } from '@elastic/eui/es/components/icon/assets/expandMini'
import { icon as __export } from '@elastic/eui/es/components/icon/assets/export'
import { icon as eye_closed } from '@elastic/eui/es/components/icon/assets/eye_closed'
import { icon as eye } from '@elastic/eui/es/components/icon/assets/eye'
import { icon as face_happy } from '@elastic/eui/es/components/icon/assets/face_happy'
import { icon as face_neutral } from '@elastic/eui/es/components/icon/assets/face_neutral'
import { icon as face_sad } from '@elastic/eui/es/components/icon/assets/face_sad'
import { icon as filter } from '@elastic/eui/es/components/icon/assets/filter'
import { icon as flag } from '@elastic/eui/es/components/icon/assets/flag'
import { icon as folder_check } from '@elastic/eui/es/components/icon/assets/folder_check'
import { icon as folder_closed } from '@elastic/eui/es/components/icon/assets/folder_closed'
import { icon as folder_exclamation } from '@elastic/eui/es/components/icon/assets/folder_exclamation'
import { icon as folder_open } from '@elastic/eui/es/components/icon/assets/folder_open'
import { icon as fold } from '@elastic/eui/es/components/icon/assets/fold'
import { icon as frameNext } from '@elastic/eui/es/components/icon/assets/frameNext'
import { icon as framePrevious } from '@elastic/eui/es/components/icon/assets/framePrevious'
import { icon as fullScreenExit } from '@elastic/eui/es/components/icon/assets/fullScreenExit'
import { icon as full_screen } from '@elastic/eui/es/components/icon/assets/full_screen'
import { icon as __function } from '@elastic/eui/es/components/icon/assets/function'
import { icon as gear } from '@elastic/eui/es/components/icon/assets/gear'
import { icon as glasses } from '@elastic/eui/es/components/icon/assets/glasses'
import { icon as globe } from '@elastic/eui/es/components/icon/assets/globe'
import { icon as grab_horizontal } from '@elastic/eui/es/components/icon/assets/grab_horizontal'
import { icon as grab } from '@elastic/eui/es/components/icon/assets/grab'
import { icon as grid } from '@elastic/eui/es/components/icon/assets/grid'
import { icon as heart } from '@elastic/eui/es/components/icon/assets/heart'
import { icon as heatmap } from '@elastic/eui/es/components/icon/assets/heatmap'
import { icon as help } from '@elastic/eui/es/components/icon/assets/help'
import { icon as home } from '@elastic/eui/es/components/icon/assets/home'
import { icon as iInCircle } from '@elastic/eui/es/components/icon/assets/iInCircle'
import { icon as image } from '@elastic/eui/es/components/icon/assets/image'
import { icon as __import } from '@elastic/eui/es/components/icon/assets/import'
import { icon as index_close } from '@elastic/eui/es/components/icon/assets/index_close'
import { icon as index_edit } from '@elastic/eui/es/components/icon/assets/index_edit'
import { icon as index_flush } from '@elastic/eui/es/components/icon/assets/index_flush'
import { icon as index_mapping } from '@elastic/eui/es/components/icon/assets/index_mapping'
import { icon as index_open } from '@elastic/eui/es/components/icon/assets/index_open'
import { icon as index_runtime } from '@elastic/eui/es/components/icon/assets/index_runtime'
import { icon as index_settings } from '@elastic/eui/es/components/icon/assets/index_settings'
import { icon as inputOutput } from '@elastic/eui/es/components/icon/assets/inputOutput'
import { icon as inspect } from '@elastic/eui/es/components/icon/assets/inspect'
import { icon as invert } from '@elastic/eui/es/components/icon/assets/invert'
import { icon as ip } from '@elastic/eui/es/components/icon/assets/ip'
import { icon as kql_field } from '@elastic/eui/es/components/icon/assets/kql_field'
import { icon as kql_function } from '@elastic/eui/es/components/icon/assets/kql_function'
import { icon as kql_operand } from '@elastic/eui/es/components/icon/assets/kql_operand'
import { icon as kql_selector } from '@elastic/eui/es/components/icon/assets/kql_selector'
import { icon as kql_value } from '@elastic/eui/es/components/icon/assets/kql_value'
import { icon as layers } from '@elastic/eui/es/components/icon/assets/layers'
import { icon as lettering } from '@elastic/eui/es/components/icon/assets/lettering'
import { icon as lineDashed } from '@elastic/eui/es/components/icon/assets/lineDashed'
import { icon as lineDotted } from '@elastic/eui/es/components/icon/assets/lineDotted'
import { icon as lineSolid } from '@elastic/eui/es/components/icon/assets/lineSolid'
import { icon as link } from '@elastic/eui/es/components/icon/assets/link'
import { icon as list_add } from '@elastic/eui/es/components/icon/assets/list_add'
import { icon as list } from '@elastic/eui/es/components/icon/assets/list'
import { icon as lock } from '@elastic/eui/es/components/icon/assets/lock'
import { icon as lockOpen } from '@elastic/eui/es/components/icon/assets/lockOpen'
import { icon as logo_aerospike } from '@elastic/eui/es/components/icon/assets/logo_aerospike'
import { icon as logo_apache } from '@elastic/eui/es/components/icon/assets/logo_apache'
import { icon as logo_app_search } from '@elastic/eui/es/components/icon/assets/logo_app_search'
import { icon as logo_aws } from '@elastic/eui/es/components/icon/assets/logo_aws'
import { icon as logo_aws_mono } from '@elastic/eui/es/components/icon/assets/logo_aws_mono'
import { icon as logo_azure } from '@elastic/eui/es/components/icon/assets/logo_azure'
import { icon as logo_azure_mono } from '@elastic/eui/es/components/icon/assets/logo_azure_mono'
import { icon as logo_beats } from '@elastic/eui/es/components/icon/assets/logo_beats'
// import { icon as logo_business_analytics } from '@elastic/eui/es/components/icon/assets/logo_business_analytics'
import { icon as logo_ceph } from '@elastic/eui/es/components/icon/assets/logo_ceph'
import { icon as logo_cloud_ece } from '@elastic/eui/es/components/icon/assets/logo_cloud_ece'
import { icon as logo_cloud } from '@elastic/eui/es/components/icon/assets/logo_cloud'
import { icon as logo_code } from '@elastic/eui/es/components/icon/assets/logo_code'
import { icon as logo_codesandbox } from '@elastic/eui/es/components/icon/assets/logo_codesandbox'
import { icon as logo_couchbase } from '@elastic/eui/es/components/icon/assets/logo_couchbase'
import { icon as logo_docker } from '@elastic/eui/es/components/icon/assets/logo_docker'
import { icon as logo_dropwizard } from '@elastic/eui/es/components/icon/assets/logo_dropwizard'
import { icon as logo_elastic } from '@elastic/eui/es/components/icon/assets/logo_elastic'
import { icon as logo_elasticsearch } from '@elastic/eui/es/components/icon/assets/logo_elasticsearch'
import { icon as logo_elastic_stack } from '@elastic/eui/es/components/icon/assets/logo_elastic_stack'
import { icon as logo_enterprise_search } from '@elastic/eui/es/components/icon/assets/logo_enterprise_search'
import { icon as logo_etcd } from '@elastic/eui/es/components/icon/assets/logo_etcd'
import { icon as logo_gcp } from '@elastic/eui/es/components/icon/assets/logo_gcp'
import { icon as logo_gcp_mono } from '@elastic/eui/es/components/icon/assets/logo_gcp_mono'
import { icon as logo_github } from '@elastic/eui/es/components/icon/assets/logo_github'
import { icon as logo_gmail } from '@elastic/eui/es/components/icon/assets/logo_gmail'
import { icon as logo_golang } from '@elastic/eui/es/components/icon/assets/logo_golang'
import { icon as logo_google_g } from '@elastic/eui/es/components/icon/assets/logo_google_g'
import { icon as logo_haproxy } from '@elastic/eui/es/components/icon/assets/logo_haproxy'
import { icon as logo_ibm } from '@elastic/eui/es/components/icon/assets/logo_ibm'
import { icon as logo_ibm_mono } from '@elastic/eui/es/components/icon/assets/logo_ibm_mono'
import { icon as logo_kafka } from '@elastic/eui/es/components/icon/assets/logo_kafka'
import { icon as logo_kibana } from '@elastic/eui/es/components/icon/assets/logo_kibana'
import { icon as logo_kubernetes } from '@elastic/eui/es/components/icon/assets/logo_kubernetes'
import { icon as logo_logging } from '@elastic/eui/es/components/icon/assets/logo_logging'
import { icon as logo_logstash } from '@elastic/eui/es/components/icon/assets/logo_logstash'
import { icon as logo_maps } from '@elastic/eui/es/components/icon/assets/logo_maps'
import { icon as logo_memcached } from '@elastic/eui/es/components/icon/assets/logo_memcached'
import { icon as logo_metrics } from '@elastic/eui/es/components/icon/assets/logo_metrics'
import { icon as logo_mongodb } from '@elastic/eui/es/components/icon/assets/logo_mongodb'
import { icon as logo_mysql } from '@elastic/eui/es/components/icon/assets/logo_mysql'
import { icon as logo_nginx } from '@elastic/eui/es/components/icon/assets/logo_nginx'
import { icon as logo_observability } from '@elastic/eui/es/components/icon/assets/logo_observability'
import { icon as logo_osquery } from '@elastic/eui/es/components/icon/assets/logo_osquery'
import { icon as logo_php } from '@elastic/eui/es/components/icon/assets/logo_php'
import { icon as logo_postgres } from '@elastic/eui/es/components/icon/assets/logo_postgres'
import { icon as logo_prometheus } from '@elastic/eui/es/components/icon/assets/logo_prometheus'
import { icon as logo_rabbitmq } from '@elastic/eui/es/components/icon/assets/logo_rabbitmq'
import { icon as logo_redis } from '@elastic/eui/es/components/icon/assets/logo_redis'
import { icon as logo_security } from '@elastic/eui/es/components/icon/assets/logo_security'
import { icon as logo_site_search } from '@elastic/eui/es/components/icon/assets/logo_site_search'
import { icon as logo_sketch } from '@elastic/eui/es/components/icon/assets/logo_sketch'
import { icon as logo_slack } from '@elastic/eui/es/components/icon/assets/logo_slack'
import { icon as logo_uptime } from '@elastic/eui/es/components/icon/assets/logo_uptime'
import { icon as logo_webhook } from '@elastic/eui/es/components/icon/assets/logo_webhook'
import { icon as logo_windows } from '@elastic/eui/es/components/icon/assets/logo_windows'
import { icon as logo_workplace_search } from '@elastic/eui/es/components/icon/assets/logo_workplace_search'
import { icon as logstash_filter } from '@elastic/eui/es/components/icon/assets/logstash_filter'
import { icon as logstash_if } from '@elastic/eui/es/components/icon/assets/logstash_if'
import { icon as logstash_input } from '@elastic/eui/es/components/icon/assets/logstash_input'
import { icon as logstash_output } from '@elastic/eui/es/components/icon/assets/logstash_output'
import { icon as logstash_queue } from '@elastic/eui/es/components/icon/assets/logstash_queue'
import { icon as magnet } from '@elastic/eui/es/components/icon/assets/magnet'
import { icon as magnifyWithExclamation } from '@elastic/eui/es/components/icon/assets/magnifyWithExclamation'
import { icon as magnifyWithMinus } from '@elastic/eui/es/components/icon/assets/magnifyWithMinus'
import { icon as magnifyWithPlus } from '@elastic/eui/es/components/icon/assets/magnifyWithPlus'
import { icon as map_marker } from '@elastic/eui/es/components/icon/assets/map_marker'
import { icon as memory } from '@elastic/eui/es/components/icon/assets/memory'
import { icon as menuDown } from '@elastic/eui/es/components/icon/assets/menuDown'
import { icon as menu } from '@elastic/eui/es/components/icon/assets/menu'
import { icon as menuLeft } from '@elastic/eui/es/components/icon/assets/menuLeft'
import { icon as menuRight } from '@elastic/eui/es/components/icon/assets/menuRight'
import { icon as menuUp } from '@elastic/eui/es/components/icon/assets/menuUp'
import { icon as merge } from '@elastic/eui/es/components/icon/assets/merge'
import { icon as minimize } from '@elastic/eui/es/components/icon/assets/minimize'
import { icon as minus_in_circle_filled } from '@elastic/eui/es/components/icon/assets/minus_in_circle_filled'
import { icon as minus_in_circle } from '@elastic/eui/es/components/icon/assets/minus_in_circle'
import { icon as minus } from '@elastic/eui/es/components/icon/assets/minus'
import { icon as ml_classification_job } from '@elastic/eui/es/components/icon/assets/ml_classification_job'
import { icon as ml_create_advanced_job } from '@elastic/eui/es/components/icon/assets/ml_create_advanced_job'
import { icon as ml_create_multi_metric_job } from '@elastic/eui/es/components/icon/assets/ml_create_multi_metric_job'
import { icon as ml_create_population_job } from '@elastic/eui/es/components/icon/assets/ml_create_population_job'
import { icon as ml_create_single_metric_job } from '@elastic/eui/es/components/icon/assets/ml_create_single_metric_job'
import { icon as ml_data_visualizer } from '@elastic/eui/es/components/icon/assets/ml_data_visualizer'
import { icon as ml_outlier_detection_job } from '@elastic/eui/es/components/icon/assets/ml_outlier_detection_job'
import { icon as ml_regression_job } from '@elastic/eui/es/components/icon/assets/ml_regression_job'
import { icon as mobile } from '@elastic/eui/es/components/icon/assets/mobile'
import { icon as moon } from '@elastic/eui/es/components/icon/assets/moon'
import { icon as nested } from '@elastic/eui/es/components/icon/assets/nested'
import { icon as node } from '@elastic/eui/es/components/icon/assets/node'
import { icon as number } from '@elastic/eui/es/components/icon/assets/number'
import { icon as offline } from '@elastic/eui/es/components/icon/assets/offline'
import { icon as online } from '@elastic/eui/es/components/icon/assets/online'
import { icon as __package } from '@elastic/eui/es/components/icon/assets/package'
import { icon as pageSelect } from '@elastic/eui/es/components/icon/assets/pageSelect'
import { icon as pagesSelect } from '@elastic/eui/es/components/icon/assets/pagesSelect'
import { icon as paint } from '@elastic/eui/es/components/icon/assets/paint'
import { icon as paper_clip } from '@elastic/eui/es/components/icon/assets/paper_clip'
import { icon as partial } from '@elastic/eui/es/components/icon/assets/partial'
import { icon as pause } from '@elastic/eui/es/components/icon/assets/pause'
import { icon as payment } from '@elastic/eui/es/components/icon/assets/payment'
import { icon as pencil } from '@elastic/eui/es/components/icon/assets/pencil'
import { icon as percent } from '@elastic/eui/es/components/icon/assets/percent'
import { icon as pin_filled } from '@elastic/eui/es/components/icon/assets/pin_filled'
import { icon as pin } from '@elastic/eui/es/components/icon/assets/pin'
import { icon as playFilled } from '@elastic/eui/es/components/icon/assets/playFilled'
import { icon as play } from '@elastic/eui/es/components/icon/assets/play'
import { icon as plus_in_circle_filled } from '@elastic/eui/es/components/icon/assets/plus_in_circle_filled'
import { icon as plus_in_circle } from '@elastic/eui/es/components/icon/assets/plus_in_circle'
import { icon as plus } from '@elastic/eui/es/components/icon/assets/plus'
import { icon as popout } from '@elastic/eui/es/components/icon/assets/popout'
import { icon as push } from '@elastic/eui/es/components/icon/assets/push'
import { icon as question_in_circle } from '@elastic/eui/es/components/icon/assets/question_in_circle'
import { icon as quote } from '@elastic/eui/es/components/icon/assets/quote'
import { icon as refresh } from '@elastic/eui/es/components/icon/assets/refresh'
import { icon as reporter } from '@elastic/eui/es/components/icon/assets/reporter'
import { icon as return_key } from '@elastic/eui/es/components/icon/assets/return_key'
import { icon as save } from '@elastic/eui/es/components/icon/assets/save'
import { icon as scale } from '@elastic/eui/es/components/icon/assets/scale'
import { icon as search } from '@elastic/eui/es/components/icon/assets/search'
import { icon as securitySignalDetected } from '@elastic/eui/es/components/icon/assets/securitySignalDetected'
import { icon as securitySignal } from '@elastic/eui/es/components/icon/assets/securitySignal'
import { icon as securitySignalResolved } from '@elastic/eui/es/components/icon/assets/securitySignalResolved'
import { icon as sessionViewer } from '@elastic/eui/es/components/icon/assets/sessionViewer'
import { icon as shard } from '@elastic/eui/es/components/icon/assets/shard'
import { icon as share } from '@elastic/eui/es/components/icon/assets/share'
import { icon as snowflake } from '@elastic/eui/es/components/icon/assets/snowflake'
import { icon as sortable } from '@elastic/eui/es/components/icon/assets/sortable'
import { icon as sort_down } from '@elastic/eui/es/components/icon/assets/sort_down'
import { icon as sortLeft } from '@elastic/eui/es/components/icon/assets/sortLeft'
import { icon as sortRight } from '@elastic/eui/es/components/icon/assets/sortRight'
import { icon as sort_up } from '@elastic/eui/es/components/icon/assets/sort_up'
import { icon as star_empty } from '@elastic/eui/es/components/icon/assets/star_empty'
import { icon as star_empty_space } from '@elastic/eui/es/components/icon/assets/star_empty_space'
import { icon as star_filled } from '@elastic/eui/es/components/icon/assets/star_filled'
import { icon as star_filled_space } from '@elastic/eui/es/components/icon/assets/star_filled_space'
import { icon as star_minus_empty } from '@elastic/eui/es/components/icon/assets/star_minus_empty'
import { icon as star_minus_filled } from '@elastic/eui/es/components/icon/assets/star_minus_filled'
import { icon as starPlusEmpty } from '@elastic/eui/es/components/icon/assets/starPlusEmpty'
import { icon as starPlusFilled } from '@elastic/eui/es/components/icon/assets/starPlusFilled'
// import { icon as stats } from '@elastic/eui/es/components/icon/assets/stats'
import { icon as stop_filled } from '@elastic/eui/es/components/icon/assets/stop_filled'
import { icon as stop } from '@elastic/eui/es/components/icon/assets/stop'
import { icon as stop_slash } from '@elastic/eui/es/components/icon/assets/stop_slash'
import { icon as storage } from '@elastic/eui/es/components/icon/assets/storage'
import { icon as string } from '@elastic/eui/es/components/icon/assets/string'
import { icon as submodule } from '@elastic/eui/es/components/icon/assets/submodule'
import { icon as sun } from '@elastic/eui/es/components/icon/assets/sun'
import { icon as swatch_input } from '@elastic/eui/es/components/icon/assets/swatch_input'
import { icon as symlink } from '@elastic/eui/es/components/icon/assets/symlink'
import { icon as table_density_compact } from '@elastic/eui/es/components/icon/assets/table_density_compact'
import { icon as table_density_expanded } from '@elastic/eui/es/components/icon/assets/table_density_expanded'
import { icon as table_density_normal } from '@elastic/eui/es/components/icon/assets/table_density_normal'
import { icon as tableOfContents } from '@elastic/eui/es/components/icon/assets/tableOfContents'
import { icon as tag } from '@elastic/eui/es/components/icon/assets/tag'
import { icon as tear } from '@elastic/eui/es/components/icon/assets/tear'
import { icon as temperature } from '@elastic/eui/es/components/icon/assets/temperature'
import { icon as timeline } from '@elastic/eui/es/components/icon/assets/timeline'
import { icon as timeRefresh } from '@elastic/eui/es/components/icon/assets/timeRefresh'
import { icon as timeslider } from '@elastic/eui/es/components/icon/assets/timeslider'
import { icon as tokenAlias } from '@elastic/eui/es/components/icon/assets/tokenAlias'
import { icon as tokenAnnotation } from '@elastic/eui/es/components/icon/assets/tokenAnnotation'
import { icon as tokenArray } from '@elastic/eui/es/components/icon/assets/tokenArray'
import { icon as tokenBinary } from '@elastic/eui/es/components/icon/assets/tokenBinary'
import { icon as tokenBoolean } from '@elastic/eui/es/components/icon/assets/tokenBoolean'
import { icon as tokenClass } from '@elastic/eui/es/components/icon/assets/tokenClass'
import { icon as tokenCompletionSuggester } from '@elastic/eui/es/components/icon/assets/tokenCompletionSuggester'
import { icon as tokenConstant } from '@elastic/eui/es/components/icon/assets/tokenConstant'
import { icon as tokenDate } from '@elastic/eui/es/components/icon/assets/tokenDate'
import { icon as tokenDenseVector } from '@elastic/eui/es/components/icon/assets/tokenDenseVector'
import { icon as tokenElement } from '@elastic/eui/es/components/icon/assets/tokenElement'
import { icon as tokenEnum } from '@elastic/eui/es/components/icon/assets/tokenEnum'
import { icon as tokenEnumMember } from '@elastic/eui/es/components/icon/assets/tokenEnumMember'
import { icon as tokenEvent } from '@elastic/eui/es/components/icon/assets/tokenEvent'
import { icon as tokenException } from '@elastic/eui/es/components/icon/assets/tokenException'
import { icon as tokenField } from '@elastic/eui/es/components/icon/assets/tokenField'
import { icon as tokenFile } from '@elastic/eui/es/components/icon/assets/tokenFile'
import { icon as tokenFlattened } from '@elastic/eui/es/components/icon/assets/tokenFlattened'
import { icon as tokenFunction } from '@elastic/eui/es/components/icon/assets/tokenFunction'
import { icon as tokenGeo } from '@elastic/eui/es/components/icon/assets/tokenGeo'
import { icon as tokenHistogram } from '@elastic/eui/es/components/icon/assets/tokenHistogram'
import { icon as tokenInterface } from '@elastic/eui/es/components/icon/assets/tokenInterface'
import { icon as tokenIP } from '@elastic/eui/es/components/icon/assets/tokenIP'
import { icon as tokenJoin } from '@elastic/eui/es/components/icon/assets/tokenJoin'
import { icon as tokenKey } from '@elastic/eui/es/components/icon/assets/tokenKey'
import { icon as tokenKeyword } from '@elastic/eui/es/components/icon/assets/tokenKeyword'
import { icon as tokenMethod } from '@elastic/eui/es/components/icon/assets/tokenMethod'
import { icon as tokenModule } from '@elastic/eui/es/components/icon/assets/tokenModule'
import { icon as tokenNamespace } from '@elastic/eui/es/components/icon/assets/tokenNamespace'
import { icon as tokenNested } from '@elastic/eui/es/components/icon/assets/tokenNested'
import { icon as tokenNull } from '@elastic/eui/es/components/icon/assets/tokenNull'
import { icon as tokenNumber } from '@elastic/eui/es/components/icon/assets/tokenNumber'
import { icon as tokenObject } from '@elastic/eui/es/components/icon/assets/tokenObject'
import { icon as tokenOperator } from '@elastic/eui/es/components/icon/assets/tokenOperator'
import { icon as tokenPackage } from '@elastic/eui/es/components/icon/assets/tokenPackage'
import { icon as tokenParameter } from '@elastic/eui/es/components/icon/assets/tokenParameter'
import { icon as tokenPercolator } from '@elastic/eui/es/components/icon/assets/tokenPercolator'
import { icon as tokenProperty } from '@elastic/eui/es/components/icon/assets/tokenProperty'
import { icon as tokenRange } from '@elastic/eui/es/components/icon/assets/tokenRange'
import { icon as tokenRankFeature } from '@elastic/eui/es/components/icon/assets/tokenRankFeature'
import { icon as tokenRankFeatures } from '@elastic/eui/es/components/icon/assets/tokenRankFeatures'
import { icon as tokenRepo } from '@elastic/eui/es/components/icon/assets/tokenRepo'
import { icon as tokenSearchType } from '@elastic/eui/es/components/icon/assets/tokenSearchType'
import { icon as tokenShape } from '@elastic/eui/es/components/icon/assets/tokenShape'
import { icon as tokenString } from '@elastic/eui/es/components/icon/assets/tokenString'
import { icon as tokenStruct } from '@elastic/eui/es/components/icon/assets/tokenStruct'
import { icon as tokenSymbol } from '@elastic/eui/es/components/icon/assets/tokenSymbol'
import { icon as tokenTag } from '@elastic/eui/es/components/icon/assets/tokenTag'
import { icon as tokenText } from '@elastic/eui/es/components/icon/assets/tokenText'
import { icon as tokenTokenCount } from '@elastic/eui/es/components/icon/assets/tokenTokenCount'
import { icon as tokenVariable } from '@elastic/eui/es/components/icon/assets/tokenVariable'
import { icon as training } from '@elastic/eui/es/components/icon/assets/training'
import { icon as trash } from '@elastic/eui/es/components/icon/assets/trash'
import { icon as unfold } from '@elastic/eui/es/components/icon/assets/unfold'
import { icon as unlink } from '@elastic/eui/es/components/icon/assets/unlink'
import { icon as userAvatar } from '@elastic/eui/es/components/icon/assets/userAvatar'
import { icon as user } from '@elastic/eui/es/components/icon/assets/user'
import { icon as users } from '@elastic/eui/es/components/icon/assets/users'
import { icon as vector } from '@elastic/eui/es/components/icon/assets/vector'
import { icon as videoPlayer } from '@elastic/eui/es/components/icon/assets/videoPlayer'
import { icon as vis_area } from '@elastic/eui/es/components/icon/assets/vis_area'
import { icon as vis_area_stacked } from '@elastic/eui/es/components/icon/assets/vis_area_stacked'
import { icon as vis_bar_horizontal } from '@elastic/eui/es/components/icon/assets/vis_bar_horizontal'
import { icon as vis_bar_horizontal_stacked } from '@elastic/eui/es/components/icon/assets/vis_bar_horizontal_stacked'
import { icon as vis_bar_vertical } from '@elastic/eui/es/components/icon/assets/vis_bar_vertical'
import { icon as vis_bar_vertical_stacked } from '@elastic/eui/es/components/icon/assets/vis_bar_vertical_stacked'
import { icon as vis_gauge } from '@elastic/eui/es/components/icon/assets/vis_gauge'
import { icon as vis_goal } from '@elastic/eui/es/components/icon/assets/vis_goal'
import { icon as vis_line } from '@elastic/eui/es/components/icon/assets/vis_line'
import { icon as vis_map_coordinate } from '@elastic/eui/es/components/icon/assets/vis_map_coordinate'
import { icon as vis_map_region } from '@elastic/eui/es/components/icon/assets/vis_map_region'
import { icon as vis_metric } from '@elastic/eui/es/components/icon/assets/vis_metric'
import { icon as vis_pie } from '@elastic/eui/es/components/icon/assets/vis_pie'
import { icon as vis_table } from '@elastic/eui/es/components/icon/assets/vis_table'
import { icon as vis_tag_cloud } from '@elastic/eui/es/components/icon/assets/vis_tag_cloud'
import { icon as vis_text } from '@elastic/eui/es/components/icon/assets/vis_text'
import { icon as vis_timelion } from '@elastic/eui/es/components/icon/assets/vis_timelion'
import { icon as vis_vega } from '@elastic/eui/es/components/icon/assets/vis_vega'
import { icon as vis_visual_builder } from '@elastic/eui/es/components/icon/assets/vis_visual_builder'
import { icon as wordWrapDisabled } from '@elastic/eui/es/components/icon/assets/wordWrapDisabled'
import { icon as wordWrap } from '@elastic/eui/es/components/icon/assets/wordWrap'
import { icon as wrench } from '@elastic/eui/es/components/icon/assets/wrench'

type IconComponentCacheType = Partial<Record<any, unknown>>

const cachedIcons: IconComponentCacheType = {
  accessibility: accessibility,
  aggregate: aggregate,
  alert: alert,
  analyzeEvent: analyze_event,
  annotation: annotation,
  apmTrace: apm_trace,
  appAddData: app_add_data,
  appAdvancedSettings: app_advanced_settings,
  appAgent: app_agent,
  appApm: app_apm,
  appAppSearch: app_app_search,
  appAuditbeat: app_auditbeat,
  appCanvas: app_canvas,
  appCases: app_cases,
  appCode: app_code,
  appConsole: app_console,
  appCrossClusterReplication: app_cross_cluster_replication,
  appDashboard: app_dashboard,
  appDevtools: app_devtools,
  appDiscover: app_discover,
  appEms: app_ems,
  appFilebeat: app_filebeat,
  appFleet: app_fleet,
  appGis: app_gis,
  appGraph: app_graph,
  appGrok: app_grok,
  appHeartbeat: app_heartbeat,
  appIndexManagement: app_index_management,
  appIndexPattern: app_index_pattern,
  appIndexRollup: app_index_rollup,
  appLens: app_lens,
  appLogs: app_logs,
  appManagement: app_management,
  appMetricbeat: app_metricbeat,
  appMetrics: app_metrics,
  appMl: app_ml,
  appMonitoring: app_monitoring,
  appNotebook: app_notebook,
  appPacketbeat: app_packetbeat,
  appPipeline: app_pipeline,
  appRecentlyViewed: app_recently_viewed,
  appReporting: app_reporting,
  appSavedObjects: app_saved_objects,
  EuiIconAppSearchProfiler: app_search_profiler,
  // app_security_analytics,
  securityApp: app_security,
  apps: apps,
  appSpaces: app_spaces,
  appSql: app_sql,
  appTimelion: app_timelion,
  appUpgradeAssistant: app_upgrade_assistant,
  appUptime: app_uptime,
  appUsersRoles: app_users_roles,
  appVisualize: app_visualize,
  appWatches: app_watches,
  appWorkplaceSearch: app_workplace_search,
  arrowDown: arrow_down,
  arrowEnd: arrowEnd,
  arrowLeft: arrow_left,
  arrowRight: arrow_right,
  arrowStart: arrowStart,
  arrowUp: arrow_up,
  asterisk: asterisk,
  beaker: beaker,
  bell: bell,
  bellSlash: bellSlash,
  bolt: bolt,
  boxesHorizontal: boxes_horizontal,
  boxesVertical: boxes_vertical,
  branch: branch,
  branchUser: branchUser,
  broom: broom,
  brush: brush,
  bug: bug,
  bullseye: bullseye,
  calendar: calendar,
  checkInCircleFilled: checkInCircleFilled,
  check: check,
  cheer: cheer,
  clock: clock,
  cloudDrizzle: cloudDrizzle,
  cloudStormy: cloudStormy,
  cloudSunny: cloudSunny,
  color: color,
  compute: compute,
  console: console,
  continuityAboveBelow: continuityAboveBelow,
  continuityAbove: continuityAbove,
  continuityBelow: continuityBelow,
  continuityWithin: continuityWithin,
  controlsHorizontal: controls_horizontal,
  controlsVertical: controls_vertical,
  copyClipboard: copy_clipboard,
  copy: copy,
  crosshairs: crosshairs,
  cross: cross,
  currency: currency,
  cut: cut,
  database: database,
  desktop: desktop,
  documentation: documentation,
  documentEdit: documentEdit,
  document: document,
  documents: documents,
  dot: dot,
  doubleArrowLeft: doubleArrowLeft,
  doubleArrowRight: doubleArrowRight,
  download: download,
  editorAlignCenter: editor_align_center,
  editorAlignLeft: editor_align_left,
  editorAlignRight: editor_align_right,
  editorBold: editor_bold,
  editorChecklist: editor_checklist,
  editorCodeBlock: editor_code_block,
  editorComment: editor_comment,
  editorDistributeHorizontal: editorDistributeHorizontal,
  editorDistributeVertical: editorDistributeVertical,
  editorHeading: editor_heading,
  editorItalic: editor_italic,
  editorItemAlignBottom: editorItemAlignBottom,
  editorItemAlignCenter: editorItemAlignCenter,
  editorItemAlignLeft: editorItemAlignLeft,
  editorItemAlignMiddle: editorItemAlignMiddle,
  editorItemAlignRight: editorItemAlignRight,
  editorItemAlignTop: editorItemAlignTop,
  editorLink: editor_link,
  editorOrderedList: editor_ordered_list,
  editorPositionBottomLeft: editorPositionBottomLeft,
  editorPositionBottomRight: editorPositionBottomRight,
  editorPositionTopLeft: editorPositionTopLeft,
  editorPositionTopRight: editorPositionTopRight,
  editorRedo: editor_redo,
  editorStrike: editor_strike,
  editorTable: editor_table,
  editorUnderline: editor_underline,
  editorUndo: editor_undo,
  editorUnorderedList: editor_unordered_list,
  email: email,
  empty: empty,
  eql: eql,
  eraser: eraser,
  exit: exit,
  expand: expand,
  expandMini: expandMini,
  export: __export,
  eyeClosed: eye_closed,
  eye: eye,
  faceHappy: face_happy,
  faceNeutral: face_neutral,
  faceSad: face_sad,
  filter: filter,
  flag: flag,
  folderCheck: folder_check,
  folderClosed: folder_closed,
  folderExclamation: folder_exclamation,
  folderOpen: folder_open,
  fold: fold,
  frameNext: frameNext,
  framePrevious: framePrevious,
  fullScreenExit: fullScreenExit,
  fullScreen: full_screen,
  function: __function,
  gear: gear,
  glasses: glasses,
  globe: globe,
  grabHorizontal: grab_horizontal,
  grab: grab,
  grid: grid,
  heart: heart,
  heatmap: heatmap,
  help: help,
  home: home,
  iInCircle: iInCircle,
  image: image,
  import: __import,
  indexClose: index_close,
  indexEdit: index_edit,
  indexFlush: index_flush,
  indexMapping: index_mapping,
  indexOpen: index_open,
  indexRuntime: index_runtime,
  indexSettings: index_settings,
  inputOutput: inputOutput,
  inspect: inspect,
  invert: invert,
  ip: ip,
  kqlField: kql_field,
  kqlFunction: kql_function,
  kqlOperand: kql_operand,
  kqlSelector: kql_selector,
  kqlValue: kql_value,
  layers: layers,
  lettering: lettering,
  lineDashed: lineDashed,
  lineDotted: lineDotted,
  lineSolid: lineSolid,
  link: link,
  listAdd: list_add,
  list: list,
  lock: lock,
  lockOpen: lockOpen,
  logoAerospike: logo_aerospike,
  logoApache: logo_apache,
  logoAppSearch: logo_app_search,
  logoAws: logo_aws,
  logoAwsMono: logo_aws_mono,
  logoAzure: logo_azure,
  logoAzureMono: logo_azure_mono,
  logoBeats: logo_beats,
  // logo_business_analytics,
  logoCeph: logo_ceph,
  logoCloudEce: logo_cloud_ece,
  logoCloud: logo_cloud,
  logoCode: logo_code,
  logoCodesandbox: logo_codesandbox,
  logoCouchbase: logo_couchbase,
  logoDocker: logo_docker,
  logoDropwizard: logo_dropwizard,
  logoElastic: logo_elastic,
  logoElasticsearch: logo_elasticsearch,
  logoElasticStack: logo_elastic_stack,
  logoEnterpriseSearch: logo_enterprise_search,
  logoEtcd: logo_etcd,
  logoGcp: logo_gcp,
  logoGcpMono: logo_gcp_mono,
  logoGithub: logo_github,
  logoGmail: logo_gmail,
  logoGolang: logo_golang,
  logoGoogleG: logo_google_g,
  logoHaproxy: logo_haproxy,
  logoIbm: logo_ibm,
  logoIbmMono: logo_ibm_mono,
  logoKafka: logo_kafka,
  logoKibana: logo_kibana,
  logoKubernetes: logo_kubernetes,
  logoLogging: logo_logging,
  logoLogstash: logo_logstash,
  logoMaps: logo_maps,
  logoMemcached: logo_memcached,
  logoMetrics: logo_metrics,
  logoMongodb: logo_mongodb,
  logoMysql: logo_mysql,
  logoNginx: logo_nginx,
  logoObservability: logo_observability,
  logoOsquery: logo_osquery,
  logoPhp: logo_php,
  logoPostgres: logo_postgres,
  logoPrometheus: logo_prometheus,
  logoRabbitmq: logo_rabbitmq,
  logoRedis: logo_redis,
  logoSecurity: logo_security,
  logoSiteSearch: logo_site_search,
  logoSketch: logo_sketch,
  logoSlack: logo_slack,
  logoUptime: logo_uptime,
  logoWebhook: logo_webhook,
  logoWindows: logo_windows,
  logoWorkplaceSearch: logo_workplace_search,
  logstashFilter: logstash_filter,
  logstashIf: logstash_if,
  logstashInput: logstash_input,
  logstashOutput: logstash_output,
  logstashQueue: logstash_queue,
  magnet: magnet,
  magnifyWithExclamation: magnifyWithExclamation,
  magnifyWithMinus: magnifyWithMinus,
  magnifyWithPlus: magnifyWithPlus,
  mapMarker: map_marker,
  memory: memory,
  menuDown: menuDown,
  menu: menu,
  menuLeft: menuLeft,
  menuRight: menuRight,
  menuUp: menuUp,
  merge: merge,
  minimize: minimize,
  minusInCircleFilled: minus_in_circle_filled,
  minusInCircle: minus_in_circle,
  minus: minus,
  mlClassificationJob: ml_classification_job,
  mlCreateAdvancedJob: ml_create_advanced_job,
  mlCreateMultiMetricJob: ml_create_multi_metric_job,
  mlCreatePopulationJob: ml_create_population_job,
  mlCreateSingleMetricJob: ml_create_single_metric_job,
  mlDataVisualizer: ml_data_visualizer,
  mlOutlierDetectionJob: ml_outlier_detection_job,
  mlRegressionJob: ml_regression_job,
  mobile: mobile,
  moon: moon,
  nested: nested,
  node: node,
  number: number,
  offline: offline,
  online: online,
  package: __package,
  pageSelect: pageSelect,
  pagesSelect: pagesSelect,
  paint: paint,
  paperClip: paper_clip,
  partial: partial,
  pause: pause,
  payment: payment,
  pencil: pencil,
  percent: percent,
  pinFilled: pin_filled,
  pin: pin,
  playFilled: playFilled,
  play: play,
  plusInCircleFilled: plus_in_circle_filled,
  plusInCircle: plus_in_circle,
  plus: plus,
  popout: popout,
  push: push,
  questionInCircle: question_in_circle,
  quote: quote,
  refresh: refresh,
  reporter: reporter,
  returnKey: return_key,
  save: save,
  scale: scale,
  search: search,
  securitySignalDetected: securitySignalDetected,
  securitySignal: securitySignal,
  securitySignalResolved: securitySignalResolved,
  sessionViewer: sessionViewer,
  shard: shard,
  share: share,
  snowflake: snowflake,
  sortable: sortable,
  sortDown: sort_down,
  sortLeft: sortLeft,
  sortRight: sortRight,
  sortUp: sort_up,
  starEmpty: star_empty,
  starEmptySpace: star_empty_space,
  starFilled: star_filled,
  starFilledSpace: star_filled_space,
  starMinusEmpty: star_minus_empty,
  starMinusFilled: star_minus_filled,
  starPlusEmpty: starPlusEmpty,
  starPlusFilled: starPlusFilled,
  // stats,
  stopFilled: stop_filled,
  stop: stop,
  stopSlash: stop_slash,
  storage: storage,
  string: string,
  submodule: submodule,
  sun: sun,
  swatchInput: swatch_input,
  symlink: symlink,
  tableDensityCompact: table_density_compact,
  tableDensityExpanded: table_density_expanded,
  tableDensityNormal: table_density_normal,
  tableOfContents: tableOfContents,
  tag: tag,
  tear: tear,
  temperature: temperature,
  timeline: timeline,
  timeRefresh: timeRefresh,
  timeslider: timeslider,
  tokenAlias: tokenAlias,
  tokenAnnotation: tokenAnnotation,
  tokenArray: tokenArray,
  tokenBinary: tokenBinary,
  tokenBoolean: tokenBoolean,
  tokenClass: tokenClass,
  tokenCompletionSuggester: tokenCompletionSuggester,
  tokenConstant: tokenConstant,
  tokenDate: tokenDate,
  tokenDenseVector: tokenDenseVector,
  tokenElement: tokenElement,
  tokenEnum: tokenEnum,
  tokenEnumAssignedUser: tokenEnumMember,
  tokenEvent: tokenEvent,
  tokenException: tokenException,
  tokenField: tokenField,
  tokenFile: tokenFile,
  tokenFlattened: tokenFlattened,
  tokenFunction: tokenFunction,
  tokenGeo: tokenGeo,
  tokenHistogram: tokenHistogram,
  tokenInterface: tokenInterface,
  tokenIp: tokenIP,
  tokenJoin: tokenJoin,
  tokenKey: tokenKey,
  tokenKeyword: tokenKeyword,
  tokenMethod: tokenMethod,
  tokenModule: tokenModule,
  tokenNamespace: tokenNamespace,
  tokenNested: tokenNested,
  tokenNull: tokenNull,
  tokenNumber: tokenNumber,
  tokenObject: tokenObject,
  tokenOperator: tokenOperator,
  tokenPackage: tokenPackage,
  tokenParameter: tokenParameter,
  tokenPercolator: tokenPercolator,
  tokenProperty: tokenProperty,
  tokenRange: tokenRange,
  tokenRankFeature: tokenRankFeature,
  tokenRankFeatures: tokenRankFeatures,
  tokenRepo: tokenRepo,
  tokenSearchType: tokenSearchType,
  tokenShape: tokenShape,
  tokenString: tokenString,
  tokenStruct: tokenStruct,
  tokenSymbol: tokenSymbol,
  tokenTag: tokenTag,
  tokenText: tokenText,
  tokenTokenCount: tokenTokenCount,
  tokenVariable: tokenVariable,
  training: training,
  trash: trash,
  unfold: unfold,
  unlink: unlink,
  userAvatar: userAvatar,
  user: user,
  users: users,
  vector: vector,
  videoPlayer: videoPlayer,
  visArea: vis_area,
  visAreaStacked: vis_area_stacked,
  visBarHorizontal: vis_bar_horizontal,
  visBarHorizontalStacked: vis_bar_horizontal_stacked,
  visBarVertical: vis_bar_vertical,
  visBarVerticalStacked: vis_bar_vertical_stacked,
  visGauge: vis_gauge,
  visGoal: vis_goal,
  visLine: vis_line,
  visMapCoordinate: vis_map_coordinate,
  visMapRegion: vis_map_region,
  visMetric: vis_metric,
  visPie: vis_pie,
  visTable: vis_table,
  visTagCloud: vis_tag_cloud,
  visText: vis_text,
  visTimelion: vis_timelion,
  visVega: vis_vega,
  visVisualBuilder: vis_visual_builder,
  wordWrapDisabled: wordWrapDisabled,
  wordWrap: wordWrap,
  wrench: wrench,
}

appendIconComponentCache(cachedIcons)
