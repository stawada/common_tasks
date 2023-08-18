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
        <div class="headline_container">
            <label class="headline">講義を選択してください</label>
        </div>
        <div class="dropdown">
            <select>
                <option value=""></option>
                <option value="math">数学</option>
                <option value="japanese">国語</option>
                <option value="english">英語</option>
            </select>
        </div>
        <div class="headline_container">
            <label class="headline">日付を選択してください</label> 
        </div>
        <div class="dropdown">
            <select>
                <option value=""></option>
                <option value="option1">7/20 11:00</option>
                <option value="option2">7/20 14:20</option>
                <option value="option3">7/22 13:30</option>
            </select>
        </div>
    </div>
</template>

<style>
select{
  -webkit-appearance: none;
  appearance: none; /* デフォルトの矢印を非表示 */
  background: #D9D9D9;
  border: 0;
  width: 304px;
  height: 62px;

  font-size: 17px;
  font-weight: 550;
  color: #494848;
  text-align: center;

  background-image: url("@/assets/pulldown.png"); /*ここに画像を入れてね */
  background-position: right 10px center;
  background-repeat: no-repeat;
  background-size: 34px 29px;
}

.headline {
  font-size: 20px;
  color: #494848;
  font-weight: 600;
}
.headline_container{
  width: 304px;
  padding-bottom: 5px;
  margin: 0 auto;
  text-align: left;
}

select::-ms-expand {
display: none; /* デフォルトの矢印を非表示(IE用) */
}

.dropdown {
  padding-bottom: 30px;
}

</style>

