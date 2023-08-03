<!--
<template>
    <div class="dropdown" style="width:100%; display:flex; justify-content: center; align-items: center;">
        <div>
            <select>
            <option value="math">数学</option>
            <option value="japanese">国語</option>
            <option value="english">英語</option>
        </select>
        </div>
        <div>
            <select>
            <option value="option1">7/20 11:00</option>
            <option value="option2">7/20 14:20</option>
            <option value="option3">7/22 13:30</option>
        </select>
        </div>
    </div>
</template>
-->

<template>
    <div class="dropdown-container">
        <div class="dropdown">
        <select>
            <option value="math">数学</option>
            <option value="japanese">国語</option>
            <option value="english">英語</option>
        </select>
        </div>
        <div class="dropdown">
        <select>
            <option value="option1">7/20 11:00</option>
            <option value="option2">7/20 14:20</option>
            <option value="option3">7/22 13:30</option>
        </select>
        </div>
    </div>
</template>

<style>
.dropdown-container {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.dropdown {
    margin: 10px;
}
</style>

