<template>
  <div class="container max-w-3xl mx-auto mt-40">
    <div class="h-60 mb-8">
      <div class="w-52 h-52 mx-auto mb-4"><vitecamp class="w-52 h-52"></vitecamp></div>
    </div>
    <div class="text-center text-md">
      <h1 class="font-serif font-bold text-4xl mb-8">{{ t('hello') }} , {{ t('welcome to') }} GoChive</h1>
      <p class="mb-10">
        <strong>GoChive</strong>
      </p>
      <p class="mb-10">

      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import vitecamp from '@/assets/svg/vitecamp.svg?component';

const { t } = useI18n();
ElMessage.info({ message: 'welcome', duration: 1 });
ElNotification({
  title: '欢迎',
  message: '',
  duration: 0,
});
</script>

<style lang="scss" scoped>
a {
  @apply text-sky-400 hover:(text-sky-600) transition-all ease-out duration-100;
}
</style>
