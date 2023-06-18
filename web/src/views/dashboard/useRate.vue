<template></template>


<script lang="ts">
import * as echarts from 'echarts/core';
import {
    TitleComponent,
    TitleComponentOption,
    CalendarComponent,
    CalendarComponentOption,
    TooltipComponent,
    TooltipComponentOption,
    VisualMapComponent,
    VisualMapComponentOption
} from 'echarts/components';
import { HeatmapChart, HeatmapSeriesOption } from 'echarts/charts';
import { CanvasRenderer } from 'echarts/renderers';

echarts.use([
    TitleComponent,
    CalendarComponent,
    TooltipComponent,
    VisualMapComponent,
    HeatmapChart,
    CanvasRenderer
]);

type EChartsOption = echarts.ComposeOption<
    | TitleComponentOption
    | CalendarComponentOption
    | TooltipComponentOption
    | VisualMapComponentOption
    | HeatmapSeriesOption
>;

var chartDom = document.getElementById('main')!;
var myChart = echarts.init(chartDom);
var option: EChartsOption;

function getVirtualData(year: string) {
    const date = +echarts.time.parse(year + '-01-01');
    const end = +echarts.time.parse(+year + 1 + '-01-01');
    const dayTime = 3600 * 24 * 1000;
    const data: [string, number][] = [];
    for (let time = date; time < end; time += dayTime) {
        data.push([
            echarts.time.format(time, '{yyyy}-{MM}-{dd}', false),
            Math.floor(Math.random() * 10000)
        ]);
    }
    return data;
}

option = {
    title: {
        top: 30,
        left: 'center',
        text: 'Daily Step Count'
    },
    tooltip: {},
    visualMap: {
        min: 0,
        max: 10000,
        type: 'piecewise',
        orient: 'horizontal',
        left: 'center',
        top: 65
    },
    calendar: {
        top: 120,
        left: 30,
        right: 30,
        cellSize: ['auto', 13],
        range: '2016',
        itemStyle: {
            borderWidth: 0.5
        },
        yearLabel: { show: false }
    },
    series: {
        type: 'heatmap',
        coordinateSystem: 'calendar',
        data: getVirtualData('2016')
    }
};

option && myChart.setOption(option);


</script>

