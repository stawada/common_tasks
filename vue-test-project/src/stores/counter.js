import { defineStore } from 'pinia';

export const useStoreCounter = defineStore('counter', {
  state: () => ({
    count: 1,
  }),
});