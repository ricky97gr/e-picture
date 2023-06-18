<script lang="ts" setup>

import LayoutHeader from './LayoutHeader.vue';
import LayoutSidebar from './LayoutSidebar.vue';
import LayoutFooter from './LayoutFooter.vue';
import '../style/layout.css'
</script>


<template>
    <div class="common-layout">
        <el-container>
            <el-header>
                <layout-header></layout-header>
            </el-header>
            <el-container>
                <el-aside>
                    <layout-sidebar></layout-sidebar>
                </el-aside>
                <el-container>
                    <el-main>
                        <router-view></router-view>
                    </el-main>
                    <el-footer>
                        <layout-footer></layout-footer>
                    </el-footer>
                </el-container>
            </el-container>
        </el-container>
    </div>
</template>

<style scoped></style>
  