<template>
    <div>
        <h3 class="title">SureLink</h3>
        <h3 class="subtitle"> You are being redirected to your link </h3>
      <div class="base-timer">
          <svg class="base-timer__svg" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
              <g class="base-timer__circle">
                  <circle class="base-timer__path-elapsed" cx="50" cy="50" r="45"></circle>
                  <path
                          :stroke-dasharray="circleDasharray"
                          class="base-timer__path-remaining"
                          :class="remainingPathColor"
                          d="
              M 50, 50
              m -45, 0
              a 45,45 0 1,0 90,0
              a 45,45 0 1,0 -90,0
            "
                  ></path>
              </g>
          </svg>
          <span class="base-timer__label">{{ formattedTimeLeft }}</span>
      </div>
    </div>
</template>

<script>
const FULL_DASH_ARRAY = 283;

const COLOR_CODES = {
    info: {
        color: "green"
    },
    warning: {
        color: "orange",
    },
    alert: {
        color: "red",
    }
};

const TIME_LIMIT = 5;

export default {
    data() {
        return {
            timePassed: 0,
            timerInterval: null
        };
    },

    computed: {
        circleDasharray() {
            return `${(this.timeFraction * FULL_DASH_ARRAY).toFixed(0)} 283`;
        },

        formattedTimeLeft() {
            const timeLeft = this.timeLeft;
            const minutes = Math.floor(timeLeft / 60);
            let seconds = timeLeft % 60;

            if (seconds < 10) {
                seconds = `0${seconds}`;
            }

            return `${minutes}:${seconds}`;
        },

        timeLeft() {
            return TIME_LIMIT - this.timePassed;
        },

        timeFraction() {
            const rawTimeFraction = this.timeLeft / TIME_LIMIT;
            return rawTimeFraction - (1 / TIME_LIMIT) * (1 - rawTimeFraction);
        },

        remainingPathColor() {
            const { info, warning, alert } = COLOR_CODES;
            if(this.timeLeft === 0)
                return alert.color;
            else if (this.timeLeft <= 2)
                return warning.color;
            else return info.color;
        }
    },

    watch: {
        timeLeft(newValue) {
            if (newValue === 0) {
                this.onTimesUp();
            }
        }
    },

    mounted() {
        this.startTimer();
    },

    methods: {
        onTimesUp() {
            clearInterval(this.timerInterval);
        },

        startTimer() {
            this.timerInterval = setInterval(() => (this.timePassed += 1), 1000);
        }
    }
};
</script>

<style scoped lang="scss">
.base-timer {
  position: relative;
  width: 300px;
  height: 300px;

  &__svg {
    transform: scaleX(-1);
  }

  &__circle {
    fill: none;
    stroke: none;
  }

  &__path-elapsed {
    stroke-width: 7px;
    stroke: grey;
  }

  &__path-remaining {
    stroke-width: 7px;
    stroke-linecap: round;
    transform: rotate(90deg);
    transform-origin: center;
    transition: 1s linear all;
    fill-rule: nonzero;
    stroke: currentColor;

    &.green {
      color: rgb(65, 184, 131);
    }

    &.orange {
      color: orange;
    }

    &.red {
      color: red;
    }
  }

  &__label {
    position: absolute;
    width: 300px;
    height: 300px;
    top: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 48px;
  }
}
</style>
