<script>
    import {task_space_url, main_space_url} from '../stores/windows'
</script>

<div class="flex-1 flex">
    <iframe src={$task_space_url} title="task space" class="flex-1 bg-blue-200">123</iframe>
    <iframe src={$main_space_url} title="main space" class="flex-[4] bg-blue-300">123</iframe>
</div>