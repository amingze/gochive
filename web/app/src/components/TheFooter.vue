<template>
  <div class="text-center">
    <el-tooltip :content="t('home')" placement="top">
      <router-link class="icon-btn mx-2" to="/">
        <i-mdi-home-search-outline class="icon-footer"/>
      </router-link>
    </el-tooltip>

    <el-tooltip :content="isDark ? t('change light') : t('change dark')" placement="top">
      <button class="icon-btn mx-2 !outline-none" @click="toggleDark()">
        <i-ph-cloud-moon-bold v-if="isDark" class="icon-footer"/>
        <i-ph-sun-horizon-bold v-else class="icon-footer"/>
      </button>
    </el-tooltip>

    <el-tooltip :content="t('change lang')" placement="top">
      <button class="icon-btn mx-2" @click="toggleLocales()">
        <i-la-language class="icon-footer"/>
      </button>
    </el-tooltip>

    <el-tooltip :content="t('method to using')" placement="top">
      <router-link class="icon-btn mx-2" to="/process">
        <i-ri-article-line class="icon-footer"/>
      </router-link>
    </el-tooltip>

    <el-tooltip :content="t('登陆')" placement="top">
      <router-link class="icon-btn mx-2" to="/login">
        <i-mdi-content-duplicate class="icon-footer"/>
      </router-link>
    </el-tooltip>
      <el-tooltip :content="t('Memo')" placement="top">
          <router-link class="icon-btn mx-2" to="/memo">
              <i-mdi-content-duplicate class="icon-footer"/>
          </router-link>
      </el-tooltip>
    <el-tooltip content="注销登陆" placement="top">
      <a class="icon-btn mx-2" @click="onLogout" target="_blank" title="GitHub">
        <i-akar-icons-github-fill class="icon-footer"/>
      </a>
    </el-tooltip>
  </div>
</template>

<script setup lang="ts">
import {isDark, toggleDark} from '@/utils/dark';
import loginApi from "@/api/modules/token";
import userStore from "@/store/user";

const {t, availableLocales, locale} = useI18n();
const toggleLocales = () => {
  const locales = availableLocales;
  locale.value = locales[(locales.indexOf(locale.value) + 1) % locales.length];
};
const onLogout=()=>{
    loginApi.deleteToken()
}
</script>

<style lang="scss">
.icon-footer {
  font-size: 1.3em;
}
</style>
