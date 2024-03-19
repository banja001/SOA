using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Encounters.API.Public;
using Explorer.Encounters.API.Dtos;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Explorer.Tours.API.Dtos;
using Explorer.Encounters.Core.UseCases;
using System.Text.Json;

namespace Explorer.API.Controllers.Administrator
{
    [Authorize(Policy = "administratorPolicy")]
    [Route("api/administrator/challenge")]
    public class ChallengeController : BaseApiController
    {
        private readonly IChallengeService _challengeService; 
        private readonly IHttpClientFactory _factory;

        public ChallengeController(IChallengeService challengeService, IHttpClientFactory factory)
        {
            _challengeService = challengeService;
            _factory = factory;
        }

        [HttpGet]
        public async Task<ActionResult<PagedResult<ChallengeDto>>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            //var result = _challengeService.GetPaged(page, pageSize);
            var client = _factory.CreateClient("encountersMicroservice");
            using HttpResponseMessage response = await client.GetAsync("challenge");
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();


            var challenges = JsonSerializer.Deserialize<List<ChallengeDto>>(jsonResponse);

            return Ok(challenges);
        }

        [HttpPost]
        public ActionResult<ChallengeDto> Create([FromBody] ChallengeDto challengeDto)
        {
            var result = _challengeService.Create(challengeDto);
            return CreateResponse(result);
        }

        [HttpPut("{id:int}")]
        public async Task<ActionResult<ChallengeDto>> Update([FromBody] ChallengeDto challengeDto)
        {
            //var result = _challengeService.Update(challengeDto);
            var client = _factory.CreateClient("encountersMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("challenge" , challengeDto);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }

        [HttpDelete("{id:int}")]
        public async Task<ActionResult> Delete(int id)
        {
            //var result = _challengeService.Delete(id);
            var client = _factory.CreateClient("encountersMicroservice");
            using HttpResponseMessage response = await client.DeleteAsync("challenge/" + id);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }
    }
}
